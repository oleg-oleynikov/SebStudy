package server

import (
	"SebStudy/adapters/primary"
	"SebStudy/adapters/util"
	"SebStudy/config"
	"SebStudy/domain/resume"
	"SebStudy/domain/resume/mapping"
	"SebStudy/infrastructure"
	"SebStudy/infrastructure/eventsourcing"
	"SebStudy/logger"
	"SebStudy/pb"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	cfg    *config.Config
	log    logger.Logger
	nc     *nats.Conn
	doneCh chan struct{}
}

func NewServer(cfg *config.Config, log logger.Logger) *server {
	return &server{cfg: cfg, log: log, doneCh: make(chan struct{})}
}

func (s *server) Run() error {
	// _, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT) // Вернуть ctx
	// defer cancel()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	log.Println(logrus.GetLevel())

	natsConn, err := nats.Connect(s.cfg.NatsUrl)
	if err != nil {
		return err
	}

	s.nc = natsConn
	defer s.nc.Close()
	s.log.Infof("Nats connected, on %s", s.cfg.NatsUrl)

	cloudeventMapper := util.NewCloudeventMapper()
	prepareCloudeventMapper(cloudeventMapper)

	eventSerde := infrastructure.NewEsEventSerde()
	jetstreamEventStore := eventsourcing.NewJetStreamEventStore(s.log, s.nc, eventSerde, "sebstudy")
	aggregateStore := eventsourcing.NewEsAggregateStore(jetstreamEventStore)

	resumeRepo := resume.NewEsResumeRepository(aggregateStore)
	resumeCmdHandlers := resume.NewResumeCommandHandlers(resumeRepo)

	cmdHandlerMap := registerCommandHandlers(resumeCmdHandlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)

	grpcServer := grpc.NewServer()
	cloudeventService := primary.NewCloudEventService(s.log, dispatcher, cloudeventMapper)
	pb.RegisterCloudEventServiceServer(grpcServer, cloudeventService)

	l, err := net.Listen("tcp", s.cfg.ServerPort)
	if err != nil {
		return err
	}

	go func() {
		s.log.Infof("gRPC server is listening on port {%s}", s.cfg.GRPC.Port)
		s.log.Error(grpcServer.Serve(l))
	}()

	<-quit
	// <-s.doneCh
	s.log.Infof("Server shutdown...")
	//
	grpcServer.GracefulStop()
	//
	return nil
}

func prepareCloudeventMapper(cloudeventMapper *util.CloudeventMapper) {
	mapping.RegisterResumeTypes(cloudeventMapper)
}

func registerCommandHandlers(cmdHandlers ...infrastructure.CommandHandlerModule) infrastructure.CommandHandlerMap {
	cmdHandlerMap := infrastructure.NewCommandHandlerMap()
	for _, handler := range cmdHandlers {
		handler.RegisterCommands(&cmdHandlerMap)
	}

	return cmdHandlerMap
}

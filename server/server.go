package server

import (
	"SebStudy/config"
	"SebStudy/domain/resume"
	"SebStudy/domain/resume/mapping"
	"SebStudy/infrastructure"
	"SebStudy/infrastructure/eventsourcing"
	"SebStudy/logger"
	"SebStudy/ports"
	"SebStudy/util"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

type server struct {
	cfg *config.Config
	log logger.Logger
	nc  *nats.Conn

	cmdDispatcher ports.CommandDispatcher
	cmdAdapter    *util.CloudEventCommandAdapter
}

func NewServer(cfg *config.Config, log logger.Logger) *server {
	return &server{cfg: cfg, log: log}
}

func (s *server) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	natsConn, err := nats.Connect(s.cfg.NatsUrl)
	if err != nil {
		return err
	}

	s.nc = natsConn
	defer s.nc.Close()
	s.log.Infof("Nats connected, on %s", s.cfg.NatsUrl)

	cloudeventMapper := util.NewCloudEventCommandAdapter()
	setupCloudeventMapper(cloudeventMapper)
	s.cmdAdapter = cloudeventMapper

	typeMapper := eventsourcing.NewTypeMapper()
	resume.RegisterResumeMappingTypes(typeMapper)

	eventSerde := eventsourcing.NewEsEventSerde(s.log, typeMapper)
	jetstreamEventStore := eventsourcing.NewJetStreamEventStore(s.log, s.nc, eventSerde, "sebstudy")
	aggregateStore := eventsourcing.NewEsAggregateStore(s.log, jetstreamEventStore)

	resumeRepo := resume.NewEsResumeRepository(aggregateStore)
	resumeCmdHandlers := resume.NewResumeCommandHandlers(resumeRepo)

	cmdHandlerMap := registerCommandHandlers(resumeCmdHandlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)
	s.cmdDispatcher = dispatcher

	closeGrpcServer, grpcServer, err := s.NewResumeGrpcServer()
	if err != nil {
		return err
	}
	defer closeGrpcServer()

	<-quit
	s.log.Infof("Server shutdown...")

	grpcServer.GracefulStop()

	return nil
}

func setupCloudeventMapper(cloudeventMapper *util.CloudEventCommandAdapter) {
	mapping.RegisterCloudeventResumeTypes(cloudeventMapper)
}

func registerCommandHandlers(cmdHandlers ...infrastructure.CommandHandlerModule) infrastructure.CommandHandlerMap {
	cmdHandlerMap := infrastructure.NewCommandHandlerMap()
	for _, handler := range cmdHandlers {
		handler.RegisterCommands(&cmdHandlerMap)
	}

	return cmdHandlerMap
}

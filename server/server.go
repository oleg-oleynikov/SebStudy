package server

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume"
	"SebStudy/internal/domain/resume/mapping"
	"SebStudy/internal/infrastructure"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/internal/infrastructure/mongodb"
	"SebStudy/internal/infrastructure/ports"
	mongoProjection "SebStudy/internal/infrastructure/projections/mongo_projection"
	"SebStudy/internal/infrastructure/repository"
	"SebStudy/internal/infrastructure/util"
	"SebStudy/logger"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	cfg         *config.Config
	log         logger.Logger
	nc          *nats.Conn
	mongoClient *mongo.Client

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
	mapping.RegisterCloudeventResumeTypes(cloudeventMapper)

	s.cmdAdapter = cloudeventMapper
	mongoDBConn, err := mongodb.NewMongoDbConn(context.Background(), *s.cfg.Mongo)
	if err != nil {
		return err
	}
	s.mongoClient = mongoDBConn

	typeMapper := eventsourcing.NewTypeMapper()
	resume.RegisterResumeMappingTypes(typeMapper)

	eventSerde := eventsourcing.NewEsEventSerde(s.log, typeMapper)
	// subManager := projections.NewSubscriptionManager(s.log, s.nc, eventSerde, nil)
	// if err = subManager.Start(context.Background()); err != nil {
	// 	s.log.Fatalf("sub manager stopped working with err: %v", err)
	// }
	mongoRepo := repository.NewMongoRepository(s.log, s.cfg, s.mongoClient)
	mongoProjection := mongoProjection.NewMongoProjection(s.log, *s.cfg, s.nc, eventSerde, mongoRepo)
	if err := mongoProjection.Start(context.Background()); err != nil {
		return err
	}

	jetstreamEventStore := eventsourcing.NewJetStreamEventStore(s.log, s.nc, eventSerde, "sebstudy")
	aggregateStore := eventsourcing.NewEsAggregateStore(s.log, jetstreamEventStore)

	resumeRepo := resume.NewEsResumeRepository(aggregateStore)
	resumeCmdHandlers := resume.NewResumeCommandHandlers(resumeRepo)

	cmdHandlerMap := infrastructure.NewCommandHandlerMap()
	resumeCmdHandlers.RegisterCommands(cmdHandlerMap)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)
	s.cmdDispatcher = dispatcher

	s.initMongoDBCollections(context.Background())

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

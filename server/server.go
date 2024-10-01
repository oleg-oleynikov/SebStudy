package server

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume"
	"SebStudy/internal/domain/resume/service"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/internal/infrastructure/mongodb"
	mongoProjection "SebStudy/internal/infrastructure/projections/mongo_projection"
	"SebStudy/internal/infrastructure/repository"
	"SebStudy/logger"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	cfg         *config.Config
	log         logger.Logger
	esdb        *esdb.Client
	mongoClient *mongo.Client
	rs          *service.ResumeService
}

func NewServer(cfg *config.Config, log logger.Logger) *server {
	return &server{cfg: cfg, log: log}
}

func (s *server) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	esdbCl, err := createESDBClient(s.cfg.EventStoreConnectionString)
	if err != nil {
		return err
	}
	s.esdb = esdbCl

	mongoDBConn, err := mongodb.NewMongoDbConn(context.Background(), *s.cfg.Mongo)
	if err != nil {
		return err
	}
	s.mongoClient = mongoDBConn

	typeMapper := eventsourcing.NewTypeMapper()
	resume.RegisterResumeMappingTypes(typeMapper)

	eventSerde := eventsourcing.NewEsEventSerde(s.log, typeMapper)

	mongoRepo := repository.NewResumeMongoRepository(s.log, s.cfg, s.mongoClient)
	mongoProjection := mongoProjection.NewMongoProjection(s.log, *s.cfg, s.esdb, eventSerde, mongoRepo)
	if err := mongoProjection.Start(context.Background()); err != nil {
		return err
	}

	eventStore := eventsourcing.NewEsEventStore(s.log, s.esdb, eventSerde, s.cfg.EventStorePrefix)
	aggregateStore := eventsourcing.NewEsAggregateStore(s.log, eventStore)

	resumeService := service.NewResumeService(s.log, s.cfg, aggregateStore, mongoRepo)
	s.rs = resumeService

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

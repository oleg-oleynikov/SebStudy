package main

import (
	"SebStudy/adapters/primary"
	"SebStudy/adapters/secondary"
	"SebStudy/adapters/util"
	"SebStudy/domain/resume"
	"SebStudy/infrastructure"

	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"

	"SebStudy/initializers"
)

const (
	// host = "localhost"
	// url = "http://localhost:8080/"
	url = "localhost:50051"
)

func main() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	eventBus, err := infrastructure.NewEventBus(nats.DefaultURL)

	if err != nil {
		log.Fatalf("failed to connect nats: %s", err)
	}

	ceMapper := util.GetCeMapperInstance()
	initializers.InitializeCeMapperHandlers()

	eventSerde := infrastructure.GetEventSerdeInstance()
	// reindexerAdapter := secondary.NewReindexerAdapter()
	writeRepo := secondary.NewPostgresAdapter()
	imageStore := infrastructure.NewImageStore("./uploads")
	eventStore := infrastructure.NewEsEventStore(eventBus, eventSerde, writeRepo, imageStore)

	// corsOptions := primary.NewCorsGrpcBuilder().WithAllowedOrigins("*").WithAllowedMethods("*").BuildHandler()
	// corsServerOption := primary.CorsToServerOptions(corsOptions)
	// ceServiceServer := primary.NewCloudEventServiceServer()
	// ceServiceServer.Run("tcp", ":50051")
	// eventStore := infrastructure.NewEsEventStore(eventBus, eventSerde, reindexerAdapter, imageStore)

	// fmt.Println(reindexerAdapter.Get("123"))

	serviceClient := infrastructure.NewCloudeventsServiceClient(url)

	ceEventSender := secondary.NewCeSenderAdapter(serviceClient, ceMapper)
	resumeRepo := resume.NewEventStoreResumeRepo(eventStore)

	resumeCmdHandlers := resume.NewHandlers(ceEventSender, resumeRepo)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap()
	cmdHandlerMap.AppendHandlers(resumeCmdHandlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)

	eventHandler := infrastructure.NewEventHandler(eventBus)

	ceAdapter := primary.NewCloudEventsAdapter(dispatcher, eventHandler, ceMapper)

	ceServiceServer := primary.NewCloudEventServiceServer(ceAdapter)
	ceServiceServer.Run("tcp", "localhost:50051")

	<-quit
	ceServiceServer.Shutdown()
}

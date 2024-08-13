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

	"github.com/joho/godotenv"

	"SebStudy/initializers"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load config: %v\n", err)
	}

	eventBus, err := infrastructure.NewEventBusNats(os.Getenv("NATS_URL"))

	if err != nil {
		log.Fatalf("failed to connect nats: %s", err)
	}

	ceMapper := util.GetCeMapperInstance()
	initializers.InitializeCeMapperHandlers()

	eventSerde := infrastructure.GetEventSerdeInstance()
	// reindexerAdapter := secondary.NewReindexerAdapter()
	writeRepo := secondary.NewPostgresAdapter()
	imageStore := infrastructure.NewImageStore(os.Getenv("IMAGE_UPLOAD_PATH"))
	eventStore := infrastructure.NewEsEventStore(eventBus, eventSerde, writeRepo, imageStore)

	// corsOptions := primary.NewCorsGrpcBuilder().WithAllowedOrigins("*").WithAllowedMethods("*").BuildHandler()
	// corsServerOption := primary.CorsToServerOptions(corsOptions)
	// ceServiceServer := primary.NewCloudEventServiceServer()
	// ceServiceServer.Run("tcp", ":50051")
	// eventStore := infrastructure.NewEsEventStore(eventBus, eventSerde, reindexerAdapter, imageStore)

	// fmt.Println(reindexerAdapter.Get("123"))

	serviceClient := infrastructure.NewCloudeventsServiceClient(os.Getenv("SERVER_URL"))

	ceEventSender := secondary.NewCeSenderAdapter(serviceClient, ceMapper)
	resumeRepo := resume.NewEventStoreResumeRepo(eventStore)

	resumeCmdHandlers := resume.NewResumeCommandHandlers(ceEventSender, resumeRepo)
	cmdHandlerMap := registerCommandHandlers(resumeCmdHandlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)

	eventHandler := infrastructure.NewEventHandler(eventBus)

	ceAdapter := primary.NewCloudEventsAdapter(dispatcher, eventHandler, ceMapper)

	ceServiceServer := primary.NewCloudEventServiceServer(ceAdapter)
	ceServiceServer.Run("tcp", os.Getenv("SERVER_URL"))

	<-quit
	ceServiceServer.Shutdown()
}

func registerCommandHandlers(cmdHandlers ...infrastructure.CommandHandlerModule) infrastructure.CommandHandlerMap {
	cmdHandlerMap := infrastructure.NewCommandHandlerMap()
	for _, handler := range cmdHandlers {
		handler.RegisterCommands(&cmdHandlerMap)
	}

	return cmdHandlerMap
}

package main

import (
	"SebStudy/adapters/primary"
	"SebStudy/adapters/secondary"
	"SebStudy/adapters/util"
	"SebStudy/domain/resume"
	"SebStudy/domain/resume/mapping"
	"SebStudy/infrastructure"
	"fmt"

	"os"
	"os/signal"
	"syscall"

	"SebStudy/infrastructure/logger"

	"SebStudy/initializers"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	initializers.LoadEnvVariables()

	serverUrl := fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT"))

	eventBus, err := infrastructure.NewEventBusNats(os.Getenv("NATS_URL"))

	if err != nil {
		logger.Logger.Fatalf("failed to connect nats: %s", err)
	}

	cloudeventMapper := util.NewCloudeventMapper()
	prepareCloudeventMapper(cloudeventMapper)

	aggregateStore := infrastructure.NewEsAggregateStore()

	eventSerde := infrastructure.GetEsEventSerdeInstance()
	writeRepo := secondary.NewPostgresAdapter()
	imageStore := infrastructure.NewImageStore(os.Getenv("IMAGE_UPLOAD_PATH"))
	eventStore := infrastructure.NewEsEventStore(eventBus, eventSerde, writeRepo, imageStore)

	// corsOptions := primary.NewCorsGrpcBuilder().WithAllowedOrigins("*").WithAllowedMethods("*").BuildHandler()
	// eventStore := infrastructure.NewEsEventStore(eventBus, eventSerde, reindexerAdapter, imageStore)

	serviceClient := infrastructure.NewCloudeventsServiceClient(serverUrl)

	ceEventSender := secondary.NewCeSenderAdapter(serviceClient, cloudeventMapper)
	resumeRepo := resume.NewEsResumeRepo(eventStore)

	resumeCmdHandlers := resume.NewResumeCommandHandlers(ceEventSender, resumeRepo)
	cmdHandlerMap := registerCommandHandlers(resumeCmdHandlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)

	eventDispatcher := infrastructure.NewEventDispatcher(eventBus)

	ceAdapter := primary.NewCloudEventsAdapter(dispatcher, eventDispatcher, cloudeventMapper)

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

func prepareCloudeventMapper(cloudeventMapper *util.CloudeventMapper) {
	mapping.RegisterResumeTypes(cloudeventMapper)
}

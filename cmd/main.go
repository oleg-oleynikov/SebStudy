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
	url = "http://localhost:8080/"
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
	writeRepo := secondary.NewPostgresAdapter()
	imageStore := infrastructure.NewImageStore("./uploads")
	eventStore := infrastructure.NewEsEventStore(eventBus, eventSerde, writeRepo, imageStore)

	ceEventSender := secondary.NewCeSenderAdapter(url, ceMapper)
	// resumeRepo := resume.NewEventStoreResumeRepo(eventStore)

	resumeCmdHandlers := resume.NewHandlers(ceEventSender, nil) // Добавить репозиторий для мракобесия ну комманд
	cmdHandlerMap := infrastructure.NewCommandHandlerMap()
	cmdHandlerMap.AppendHandlers(resumeCmdHandlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)

	eventHandler := infrastructure.NewEventHandler(eventBus, eventStore)

	ceAdapter := primary.NewCloudEventsAdapter(dispatcher, eventHandler, ceMapper)

	ceAdapter.Run()

	<-quit
}

package main

import (
	"SebStudy/adapters/primary"
	"SebStudy/adapters/secondary"
	"SebStudy/adapters/util"
	"SebStudy/adapters/util/ce_mapper_func"
	"SebStudy/domain/resume"
	"SebStudy/infrastructure"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

const (
	host = "localhost"
	port = 8080
	url  = "http://localhost:8080/"
)

func main() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	eventBus, err := infrastructure.NewEventBus(nats.DefaultURL)

	if err != nil {
		log.Fatalf("failed to connect nats: %s", err)
	}

	// eventBus.Subscribe("hello", func(msg *nats.Msg) {
	// 	fmt.Printf("Received event: %s\n", string(msg.Data))
	// })

	// eventBus.Publish("hello", []byte("Hello"))

	ceMapper := util.GetCeMapperInstance()
	ce_mapper_func.InitializeMapHandler()

	ceEventSender := secondary.NewCeSenderAdapter(url, ceMapper)

	handlers := resume.NewHandlers(ceEventSender)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap()
	cmdHandlerMap.AppendHandlers(handlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)
	eventHandlerMap := infrastructure.NewEventHandlerMap()
	eventHandler := infrastructure.NewEventHandler(eventBus, eventHandlerMap)

	ceAdapter := primary.NewCloudEventsAdapter(dispatcher, eventHandler, ceMapper, port)

	ceAdapter.Run()

	<-quit
}

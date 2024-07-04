package main

import (
	"SebStudy/adapters/primary"
	"SebStudy/adapters/secondary"
	"SebStudy/adapters/util"
	"SebStudy/domain/resume"
	"SebStudy/infrastructure"
	"os"
	"os/signal"
	"syscall"
)

const (
	host = "localhost"
	port = 8080
	url  = "http://localhost:8080/"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// _, err := infrastructure.NewEventBus(nats.DefaultURL)

	// if err != nil {
	// 	log.Fatalf("failed to connect nats: %s", err)
	// }

	// eventBus.Subscribe("hello", func(msg *nats.Msg) {
	// 	fmt.Printf("Received event: %s\n", string(msg.Data))
	// })

	// eventBus.Publish("hello", []byte("Hello"))

	ceMapper := util.NewCeMapper()

	ceEventSender := secondary.NewCeSenderAdapter(url, ceMapper)

	handlers := resume.NewHandlers(nil, ceEventSender)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap()
	cmdHandlerMap.AppendHandlers(handlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)

	ceAdapter := primary.NewCloudEventsAdapter(dispatcher, ceMapper, port)

	ceAdapter.Run()

	<-quit
}

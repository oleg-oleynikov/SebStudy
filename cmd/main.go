package main

import (
	"SebStudy/adapters/primary"
	"SebStudy/domain/resume"
	"SebStudy/infrastructure"
	"os"
	"os/signal"
	"syscall"
)

const (
	port = 8080
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

	handlers := resume.NewHandlers(nil)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap()
	cmdHandlerMap.AppendHandlers(handlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)

	ceMapper := primary.NewCeMapper()

	ceAdapter := primary.NewCloudEventsAdapter(dispatcher, ceMapper, port)

	ceAdapter.Run()

	<-quit
}

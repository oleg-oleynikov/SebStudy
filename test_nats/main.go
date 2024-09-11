package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	streamName := "sebstudy_resume_0191dc80b36a7a67a1c6332cf1361c83"

	// Попробуйте создать или обновить поток
	_, err = js.AddStream(&nats.StreamConfig{
		Name:      streamName,
		Subjects:  []string{streamName + ".>"},
		Retention: nats.LimitsPolicy,
		Storage:   nats.FileStorage,
	})
	if err != nil {
		log.Fatalf("Failed to create or update stream: %v", err)
	}

	fmt.Println("Stream created or updated successfully")
}

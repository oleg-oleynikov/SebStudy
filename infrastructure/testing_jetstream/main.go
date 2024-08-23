package main

import (
	"SebStudy/domain/resume/events"
	"SebStudy/domain/resume/values"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
)

func main() {
	// logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)

	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		logrus.Fatalf("Failed to connect nats: %v", err)
	}

	if nc == nil || !nc.IsConnected() {
		logrus.Fatalf("Fucking fuck")
	}

	// defer nc.Drain()

	js, err := jetstream.New(nc)
	if err != nil {
		logrus.Fatalf("Failed to get jetstream: %v", err)
	}

	cfgStream := jetstream.StreamConfig{
		Name:      "DOMAIN_EVENTS",
		Retention: jetstream.WorkQueuePolicy, // Не ебу что делает но еще есть 2 штуки, эта вроде недолго хранит в памяти msg,
		// или до первого потребителя что ли
		Subjects: []string{"events.>"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = js.CreateStream(ctx, cfgStream)

	if err != nil {
		logrus.Fatalf("Что т не так при создании потока %v", err)
	}

	eventUuid, _ := uuid.NewV7()

	event := events.ResumeCreated{
		ResumeId:    values.NewResumeId(eventUuid.String()),
		FirstName:   values.FirstName{FirstName: "zombi"},
		MiddleName:  values.MiddleName{MiddleName: "fucking"},
		LastName:    values.LastName{LastName: "nigger"},
		PhoneNumber: values.PhoneNumber{PhoneNumber: "79985342810"},
		Education:   values.Education{Education: "PTY"},
		AboutMe:     values.AboutMe{AboutMe: "I am guy"},
		Skills: values.Skills{
			Skills: []values.Skill{
				{Skill: "suck dick"},
			},
		},
		Photo:         values.Photo{Url: "", Photo: []byte{0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0}},
		Directions:    values.Directions{},
		AboutProjects: values.AboutProjects{AboutProjects: "about projects"},
		Portfolio:     values.Portfolio{Portfolio: "portfolio"},
		StudentGroup:  values.StudentGroup{StudentGroup: "IS-32"},
		CreatedAt:     time.Now(),
	}

	eventBytes, err := json.Marshal(event)
	logrus.Printf("При отправке байтики: %v\n", eventBytes)

	if err != nil {
		logrus.Fatalf("АНЛАК при кодировании ивента в json: %v", err)
	}

	consCfg := jetstream.ConsumerConfig{
		Durable:   "domain_event_consumer",
		AckPolicy: jetstream.AckExplicitPolicy,
	}

	consumer, err := js.CreateOrUpdateConsumer(ctx, "DOMAIN_EVENTS", consCfg)
	if err != nil {
		logrus.Fatalf("Ошибка при создании/обновлении потребителя: %v", err)
	}

	js.PublishAsync("events1.123133214", eventBytes)

	msgs, err := consumer.Fetch(10)
	if err != nil {
		logrus.Println("FUCK")
	}
	for msg := range msgs.Messages() {
		logrus.Debugf("Блядство, а это ивент: %v", msg.Data())

		logrus.Debugf("Res: %v", compare(eventBytes, msg.Data()))

		event := events.ResumeCreated{}
		if err := json.Unmarshal(msg.Data(), &event); err != nil {
			logrus.Debugf("Ошибочка при десере: %v", err)
		}

		log.Println("CYKA", event)

		err := msg.Ack()
		if err != nil {
			logrus.Printf("Failed to ack message: %v", err)
		}
	}
}

func compare(arr1 []byte, arr2 []byte) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

// type EsEventStore struct {
// 	nc      *nats.Conn
// 	prefix  string
// 	esSerde infrastructure.EventSerde
// }

// func AppendEvents(events []interface{}) {

// }

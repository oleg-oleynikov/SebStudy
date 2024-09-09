package main

import (
	"SebStudy/config"
	"SebStudy/domain/resume"
	"SebStudy/domain/resume/events"
	"SebStudy/domain/resume/values"
	"SebStudy/infrastructure"
	"SebStudy/infrastructure/eventsourcing"
	"SebStudy/logger"

	// "log"

	// "log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.InitConfig()
	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()

	nc, err := nats.Connect(nats.DefaultURL)

	serde := infrastructure.NewEsEventSerde()
	eventStore := eventsourcing.NewJetStreamEventStore(appLogger, nc, serde, "sebstudy")

	aggregateStore := eventsourcing.NewEsAggregateStore(appLogger, eventStore)

	if err != nil {
		appLogger.Fatalf("Failed to connect nats: %v", err)
	}

	if nc == nil || !nc.IsConnected() {
		appLogger.Fatalf("Fucking fuck")
	}

	eventUuid, _ := uuid.NewV7()

	event := events.ResumeCreated{
		ResumeId:    values.NewResumeId(eventUuid.String()),
		FirstName:   values.FirstName{FirstName: "vitas"},
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

	resume := resume.NewResume()
	resume.Raise(event)

	md := infrastructure.CommandMetadata{AggregateId: "1234"}

	if err := aggregateStore.Save(resume, md); err != nil {
		logrus.Debugf("Failed to save aggregate: %v", err)
		return
	}
}

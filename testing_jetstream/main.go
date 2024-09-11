package main

import (
	"SebStudy/config"
	"SebStudy/domain/resume"
	"SebStudy/domain/resume/events"
	"SebStudy/domain/resume/values"
	"SebStudy/infrastructure"
	"SebStudy/infrastructure/eventsourcing"
	"SebStudy/logger"
	"fmt"

	// "log"

	// "log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func GenerateUuidWithoutDashes() string {
	u, _ := uuid.NewV7()
	bytes, _ := u.MarshalBinary()

	uuidString := fmt.Sprintf("%x", bytes)

	return uuidString
}

func main() {
	cfg := config.InitConfig()
	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()

	nc, err := nats.Connect(nats.DefaultURL)

	typeMapper := eventsourcing.NewTypeMapper()
	resume.RegisterResumeMappingTypes(typeMapper)

	serde := eventsourcing.NewEsEventSerde(appLogger, typeMapper)
	eventStore := eventsourcing.NewJetStreamEventStore(appLogger, nc, serde, "sebstudy")

	aggregateStore := eventsourcing.NewEsAggregateStore(appLogger, eventStore)

	if err != nil {
		appLogger.Fatalf("Failed to connect nats: %v", err)
	}

	if nc == nil || !nc.IsConnected() {
		appLogger.Fatalf("nats is disconected")
	}

	resumeUuid := GenerateUuidWithoutDashes()

	event := events.ResumeCreated{
		ResumeId:    values.NewResumeId(resumeUuid),
		FirstName:   values.FirstName{FirstName: "vitas"},
		MiddleName:  values.MiddleName{MiddleName: "fucking"},
		LastName:    values.LastName{LastName: "nigger"},
		PhoneNumber: values.PhoneNumber{PhoneNumber: "79985342810"},
		Education:   values.Education{Education: "PTY"},
		AboutMe:     values.AboutMe{AboutMe: "I am guy"},
		Skills: values.Skills{
			Skills: []values.Skill{
				{Skill: "suck dick"},
				{Skill: "work"},
			},
		},
		Photo:         values.Photo{Url: "", Photo: []byte{0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0}},
		Direction:     values.Direction{},
		AboutProjects: values.AboutProjects{AboutProjects: "about projects"},
		Portfolio:     values.Portfolio{Portfolio: "portfolio"},
		CreatedAt:     time.Now(),
	}

	resume1 := resume.NewResume()
	resume1.Raise(event)

	md := infrastructure.CommandMetadata{AggregateId: "1234"}

	if err := aggregateStore.Save(resume1, md); err != nil {
		logrus.Debugf("Failed to save aggregate: %v", err)
		return
	}

	loadingResume := resume.NewResume()
	if err := aggregateStore.Load(resume1.GetId(), loadingResume); err != nil {
		appLogger.Fatalf("ОШИБКА ПРИ ЗАГРУЗКЕ АГРЕГАТА: %v", err)
	}

	appLogger.Debugf("After loading aggregate: %v", loadingResume)

}

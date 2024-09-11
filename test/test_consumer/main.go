package main

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/logger"
	"context"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	cfg := config.InitConfig()
	log := logger.NewAppLogger(cfg.Logger)
	log.InitLogger()

	typeMapper := eventsourcing.NewTypeMapper()
	resume.RegisterResumeMappingTypes(typeMapper)

	serde := eventsourcing.NewEsEventSerde(log, typeMapper)

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("nats is not connected: %v", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf("Fail to get jetstream from nats conn: %v", err)
	}

	// stream - sebstudy_Resume_0191dfe26b7f7daabad21bd475c4bd6f
	// event - sebstudy_Resume_0191dfe26b7f7daabad21bd475c4bd6f.resumeCreated_0191dfe26b997748bbfa781d8909d154
	stream, err := js.Stream(context.Background(), "sebstudy_Resume_0191dfe26b7f7daabad21bd475c4bd6f")
	if err != nil {
		log.Fatalf("Fucking fuck")
	}

	info, err := stream.Info(context.TODO())
	if err != nil {
		log.Fatalf("CYKA BLYAT")
	}

	log.Debugf("УРААААААА info: %v", info.State.Msgs)

	consumerCfg := jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverAllPolicy,
		AckPolicy:     jetstream.AckExplicitPolicy,
		Durable:       "my_durable_consumer",
	}

	consumer, err := js.CreateOrUpdateConsumer(context.Background(), "sebstudy_Resume_0191dfe26b7f7daabad21bd475c4bd6f", consumerCfg)
	if err != nil {
		log.Fatalf("Failed to get consumer: %v", err)
	}

	batch, err := consumer.Fetch(1, jetstream.FetchMaxWait(3*time.Second))

	if err != nil {
		log.Fatalf("Failed to get batch: %v", err)
	}

	for i := range batch.Messages() {
		log.Debugf("%v", i)
		event, md, err := serde.Deserialize(i)
		if err != nil {
			log.Debugf("Ошибка какая то")
		}

		log.Debugf("Сам ивент: %v", event)
		log.Debugf("Метаданные: %v", md)
		i.Ack()
	}
}

package mongo_projection

import (
	"context"
	"fmt"
	"reflect"
	"resume-server/config"
	"resume-server/internal/domain/resume/events"
	"resume-server/internal/infrastructure"
	"resume-server/internal/infrastructure/eventsourcing"
	"resume-server/internal/infrastructure/repository"
	"resume-server/logger"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/esdb"
)

type mongoProjection struct {
	log        logger.Logger
	cfg        config.Config
	esdb       *esdb.Client
	serde      eventsourcing.EventSerde
	resumeRepo repository.ResumeRepository
}

func NewMongoProjection(log logger.Logger, cfg config.Config, esdb *esdb.Client, serde eventsourcing.EventSerde, resumeRepo repository.ResumeRepository) *mongoProjection {
	return &mongoProjection{
		log:        log,
		cfg:        cfg,
		esdb:       esdb,
		serde:      serde,
		resumeRepo: resumeRepo,
	}
}
func (m *mongoProjection) Start(ctx context.Context) error {
	err := m.CreatePersistentSubscription(ctx, m.cfg.EventStoreGroupName, m.cfg.EventStorePrefix)
	if err != nil {
		return err
	}

	sub, err := m.ConnectToPersistentSubscription(ctx, m.cfg.EventStoreGroupName)
	if err != nil {
		return err
	}

	go func() {
		for {
			s := sub.Recv()
			if s.EventAppeared == nil {
				continue
			}

			eventType := s.EventAppeared.Event.EventType

			event, md, err := m.serde.Deserialize(s.EventAppeared)
			if err != nil {
				m.log.Debugf("failed to deserialize %s: %v", eventType, err)
				continue
			}

			if err := m.When(ctx, event, md); err != nil {
				m.log.Debugf("When error: %v", err)
			}

			sub.Ack(s.EventAppeared)

			if s.SubscriptionDropped != nil {
				panic(s.SubscriptionDropped.Error)
			}
		}
	}()

	return nil
}

func (m *mongoProjection) When(ctx context.Context, event interface{}, md *infrastructure.EventMetadata) error {
	switch event := event.(type) {
	case events.ResumeCreated:
		return m.onResumeCreate(ctx, event, md)
	case events.ResumeChanged:
		return m.onResumeChanged(ctx, event, md)
	default:
		m.log.Debugf("(mongoProjection) [When unknown EventType] eventType: {%s}", reflect.TypeOf(event).Name())
		return fmt.Errorf("invalid event type")
	}
}

func (m *mongoProjection) CreatePersistentSubscription(ctx context.Context, groupName string, streamNamePrefix string) error {
	opts := esdb.PersistentAllSubscriptionOptions{
		From: esdb.EndPosition,
		Filter: &esdb.SubscriptionFilter{
			Type:     esdb.StreamFilterType,
			Prefixes: []string{streamNamePrefix},
		},
	}

	err := m.esdb.CreatePersistentSubscriptionAll(ctx, groupName, opts)
	if err != nil {
		if strings.Contains(err.Error(), "AlreadyExists") {
			return nil
		}

		m.log.Debugf("(mongoProjection) error create persistent sub: %v", err)
		return err
	}

	return nil
}

func (m *mongoProjection) ConnectToPersistentSubscription(ctx context.Context, groupName string) (*esdb.PersistentSubscription, error) {
	opts := esdb.ConnectToPersistentSubscriptionOptions{}

	sub, err := m.esdb.ConnectToPersistentSubscription(ctx, "$all", groupName, opts)
	if err != nil {
		if subscriptionError, ok := err.(*esdb.PersistentSubscriptionError); !ok || ok && (subscriptionError.Code != 6) {
			return nil, err
		}

		return nil, err
	}

	return sub, nil
}

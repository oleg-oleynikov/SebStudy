package util

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	v1 "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/options/grpc/protobuf/v1"
)

type EventType string

type CloudeventToEvent func(ctx context.Context, cloudEvent *v1.CloudEvent) (interface{}, error)
type EventToCloudevent func(eventType, source string, event interface{}) (*v1.CloudEvent, error)

const (
	EVENT   EventType = "DOMAIN_EVENT"
	COMMAND EventType = "COMMAND_EVENT"
)

type CloudeventMapper struct {
	cloudeventToEvent map[string]CloudeventToEvent
	eventToCloudEvent map[reflect.Type]EventToCloudevent
	mappingTypes      map[string]EventType
}

func NewCloudeventMapper() *CloudeventMapper {
	return &CloudeventMapper{
		cloudeventToEvent: make(map[string]CloudeventToEvent, 0),
		eventToCloudEvent: make(map[reflect.Type]EventToCloudevent, 0),
		mappingTypes:      make(map[string]EventType, 0),
	}
}

func (m *CloudeventMapper) GetCloudeventToEvent(cloudeventType string) (CloudeventToEvent, error) {
	mapper, ok := m.cloudeventToEvent[cloudeventType]
	if !ok {
		return nil, fmt.Errorf("cannot find mapper for %s", cloudeventType)
	}
	return mapper, nil
}

func (m *CloudeventMapper) GetEventToCloudevent(eventType reflect.Type) (EventToCloudevent, error) {
	mapper, ok := m.eventToCloudEvent[eventType]
	if !ok {
		return nil, fmt.Errorf("cannot find mapper %s", eventType)
	}
	return mapper, nil
}

func (m *CloudeventMapper) MapEvent(eventType reflect.Type, cloudeventType string, eType EventType, toEvent CloudeventToEvent, toCloudevent EventToCloudevent) error {
	if cloudeventType == "" {
		return fmt.Errorf("need ceType")
	}

	if _, exists := m.cloudeventToEvent[cloudeventType]; exists {
		return fmt.Errorf("already exist %s", cloudeventType)
	}

	if _, exists := m.eventToCloudEvent[eventType]; exists {
		return fmt.Errorf("already exist %s", eventType)
	}

	m.cloudeventToEvent[cloudeventType] = toEvent
	m.eventToCloudEvent[eventType] = toCloudevent
	m.mappingTypes[cloudeventType] = eType

	return nil
}

func (m *CloudeventMapper) MapCommand(cloudeventType string, eType EventType, toEvent CloudeventToEvent) error {
	if cloudeventType == "" {
		return fmt.Errorf("need ceType")
	}

	if _, exists := m.cloudeventToEvent[cloudeventType]; exists {
		return fmt.Errorf("already exist %s", cloudeventType)
	}

	m.cloudeventToEvent[cloudeventType] = toEvent
	m.mappingTypes[cloudeventType] = eType

	return nil
}

func (m *CloudeventMapper) IsEvent(cloudeventType string) bool {
	ceType, ok := m.mappingTypes[cloudeventType]
	if !ok {
		return false
	}
	return ceType == EVENT
}

func (m *CloudeventMapper) IsCommand(cloudeventType string) bool {
	ceType, ok := m.mappingTypes[cloudeventType]
	if !ok {
		return false
	}
	return ceType == COMMAND
}

func GetValueType(t interface{}) reflect.Type {
	v := reflect.ValueOf(t)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Type()
}

func InitCloudEvent(eventType, source string, mes proto.Message) (*v1.CloudEvent, error) {
	eventId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	protoEvent, err := anypb.New(mes)
	if err != nil {
		return nil, err
	}

	return &v1.CloudEvent{
		Id:          eventId.String(),
		Source:      source,
		SpecVersion: "1.0",
		Type:        eventType,
		Data: &v1.CloudEvent_ProtoData{
			ProtoData: protoEvent,
		},
	}, nil
}

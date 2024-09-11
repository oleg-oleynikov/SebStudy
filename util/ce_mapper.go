package util

import (
	"SebStudy/pb"
	"context"
	"fmt"
	"reflect"
)

type CloudeventToEvent func(ctx context.Context, cloudEvent *pb.CloudEvent) (interface{}, error)
type EventToCloudevent func(eventType, source string, event interface{}) (*pb.CloudEvent, error)

type CloudEventCommandAdapter struct {
	cloudeventToEvent map[string]CloudeventToEvent
	eventToCloudEvent map[reflect.Type]EventToCloudevent
}

func NewCloudEventCommandAdapter() *CloudEventCommandAdapter {
	return &CloudEventCommandAdapter{
		cloudeventToEvent: make(map[string]CloudeventToEvent, 0),
		eventToCloudEvent: make(map[reflect.Type]EventToCloudevent, 0),
	}
}

func (m *CloudEventCommandAdapter) GetCloudeventToEvent(cloudeventType string) (CloudeventToEvent, error) {
	mapper, ok := m.cloudeventToEvent[cloudeventType]
	if !ok {
		return nil, fmt.Errorf("cannot find mapper for %s", cloudeventType)
	}
	return mapper, nil
}

func (m *CloudEventCommandAdapter) GetEventToCloudevent(eventType reflect.Type) (EventToCloudevent, error) {
	mapper, ok := m.eventToCloudEvent[eventType]
	if !ok {
		return nil, fmt.Errorf("cannot find mapper %s", eventType)
	}
	return mapper, nil
}

func (m *CloudEventCommandAdapter) MapCloudevent(cloudeventType string, toEvent CloudeventToEvent) error {
	if cloudeventType == "" {
		return fmt.Errorf("need ceType")
	}

	if _, exists := m.cloudeventToEvent[cloudeventType]; exists {
		return fmt.Errorf("already exist %s", cloudeventType)
	}

	m.cloudeventToEvent[cloudeventType] = toEvent

	return nil
}

func GetValueType(t interface{}) reflect.Type {
	v := reflect.ValueOf(t)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Type()
}

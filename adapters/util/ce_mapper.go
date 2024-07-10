package util

import (
	"context"
	"encoding/base64"
	"fmt"
	"sync"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"google.golang.org/protobuf/proto"
)

type EventType string

type CeToEvent func(ctx context.Context, cloudEvent cloudevents.Event) (interface{}, error)
type EventToCe func(eventType, source string, e interface{}) (cloudevents.Event, error)

const (
	EVENT   EventType = "event"
	COMMAND EventType = "command"
)

type CeMapper struct {
	CeToEvent  map[string]CeToEvent
	EventToCe  map[string]EventToCe
	eventTypes map[string]EventType
}

var (
	instance1 *CeMapper
	once1     sync.Once
)

func GetCeMapperInstance() *CeMapper {
	once1.Do(func() {
		c := &CeMapper{}
		c.CeToEvent = make(map[string]CeToEvent, 0)
		c.EventToCe = make(map[string]EventToCe, 0)
		c.eventTypes = make(map[string]EventType, 0)

		instance1 = c
	})

	return instance1
}

func (cm *CeMapper) MapToEvent(ctx context.Context, c cloudevents.Event) (interface{}, error) {
	handler, err := cm.Get(c.Type())
	if err != nil {
		return nil, err
	}

	cmd, err := handler(ctx, c)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cm *CeMapper) MapToCloudEvent(e interface{}, eventType, source string) (cloudevents.Event, error) {
	handler, err := cm.GetToCe(eventType)
	if err != nil {
		return cloudevents.Event{}, err
	}

	cloudEvent, err := handler(eventType, source, e)
	if err != nil {
		return cloudevents.Event{}, err
	}

	return cloudEvent, nil
}

func (cm *CeMapper) Get(t string) (CeToEvent, error) {
	if h, ex := cm.CeToEvent[t]; ex {
		return h, nil
	}

	return nil, fmt.Errorf("handler for {%v} type doesnt exist", t)
}

func (cm *CeMapper) GetToCe(t string) (EventToCe, error) {
	if h, ex := cm.EventToCe[t]; ex {
		return h, nil
	}

	return nil, fmt.Errorf("handler for {%v} type doesnt exist", t)
}

func (c *CeMapper) register(ceType string, f CeToEvent, typeEvent EventType) {
	c.CeToEvent[ceType] = f
	c.eventTypes[ceType] = typeEvent
}

func (c *CeMapper) RegisterEvent(ceType string, e CeToEvent, ce EventToCe) {
	c.register(ceType, e, EVENT)
	c.EventToCe[ceType] = ce
}

func (c *CeMapper) RegisterCommand(ceType string, e CeToEvent) {
	c.register(ceType, e, COMMAND)
}

func (c *CeMapper) IsCommand(ceType string) bool {
	typeOfEvent, _ := c.GetEventType(ceType)
	return typeOfEvent == COMMAND
}

func (c *CeMapper) IsEvent(ceType string) bool {
	typeOfEvent, _ := c.GetEventType(ceType)
	return typeOfEvent == EVENT
}

func (c *CeMapper) GetEventType(ceType string) (EventType, error) {
	typeEvent, ok := c.eventTypes[ceType]
	if ok {
		return typeEvent, nil
	}
	return typeEvent, fmt.Errorf("type %s does not exist", ceType)
}

func InitCloudEvent(eventType, source string, mes proto.Message) cloudevents.Event {
	cloudEvent := cloudevents.Event{}
	cloudEvent.SetSpecVersion("1.0")
	cloudEvent.SetType(eventType)
	cloudEvent.SetSource(source)
	b, _ := proto.Marshal(mes)
	var protoBytes []byte = make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(protoBytes, b)
	cloudEvent.SetData("application/protobuf", protoBytes)

	return cloudEvent
}

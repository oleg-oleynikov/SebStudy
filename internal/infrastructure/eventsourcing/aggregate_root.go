package eventsourcing

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type AggregateRoot interface {
	Load(events []interface{})
	ClearChanges()
	GetChanges() []interface{}
	GetId() string
	GetVersion() int
}

type AggregateRootBase struct {
	AggregateRoot

	Id       string
	version  int
	changes  []interface{}
	handlers map[reflect.Type]func(interface{})
}

func (a *AggregateRootBase) GenerateUuidWithoutDashes() string {
	u, _ := uuid.NewV7()
	bytes, _ := u.MarshalBinary()

	uuidString := fmt.Sprintf("%x", bytes)

	return uuidString
}

func NewAggregateRootBase() AggregateRootBase {
	return AggregateRootBase{
		version:  -1,
		changes:  make([]interface{}, 0),
		handlers: make(map[reflect.Type]func(interface{}), 0),
	}
}

func getValueType(t interface{}) reflect.Type {
	v := reflect.ValueOf(t)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Type()
}

func GetStreamName(a AggregateRoot) string {
	return GetStreamNameWithId(a, a.GetId())
}

func GetStreamNameWithId(a AggregateRoot, id string) string {
	return getValueType(a).Name() + "_" + id
}

func (a *AggregateRootBase) Register(event interface{}, handler func(interface{})) {
	a.handlers[getValueType(event)] = handler
}

func (a *AggregateRootBase) Load(events []interface{}) {
	for _, event := range events {
		a.Raise(event)
		a.version++
	}
}

func (a *AggregateRootBase) Raise(event interface{}) {
	if handler, exists := a.handlers[getValueType(event)]; exists {
		handler(event)
		a.changes = append(a.changes, event)
	}
}

func (a *AggregateRootBase) ClearChanges() {
	a.changes = []interface{}{}
}

func (a *AggregateRootBase) GetChanges() []interface{} {
	return a.changes
}

func (a *AggregateRootBase) GetId() string {
	return a.Id
}

func (a *AggregateRootBase) GetVersion() int {
	return a.version
}

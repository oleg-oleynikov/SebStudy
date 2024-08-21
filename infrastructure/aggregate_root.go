package infrastructure

import (
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

func NewAggregateRootBase() AggregateRootBase {
	uuid, _ := uuid.NewV7()
	return AggregateRootBase{
		Id:       uuid.String(),
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
	return getValueType(a).String() + "-" + id
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

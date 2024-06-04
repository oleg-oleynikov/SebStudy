package eventsourcing

import "reflect"

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

func NewAggregateRootBase() *AggregateRootBase {
	return &AggregateRootBase{
		version:  -1,
		changes:  make([]interface{}, 0),
		handlers: make(map[reflect.Type]func(interface{})),
	}
}

func getValueType(t interface{}) reflect.Type {
	v := reflect.ValueOf(t)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Type()
}

func (a *AggregateRootBase) Register(event interface{}, handler func(interface{})) {
	a.handlers[getValueType(event)] = handler
}

func (a *AggregateRootBase) Load(events []interface{}) {
	for _, event := range events {
		a.Apply(event)
	}
}

func (a *AggregateRootBase) Apply(event interface{}) {
	if handler, exists := a.handlers[getValueType(event)]; exists {
		handler(event)
		a.changes = append(a.changes, event)
		a.version++
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

package infrastructure

import (
	"fmt"
	"reflect"
)

type CommandHandle func(Command, CommandMetadata) error

type CommandHandlerModule interface {
	RegisterCommands(cmdHandlerMap *CommandHandlerMap)
}

type CommandHandlerMap struct {
	handlers map[reflect.Type]CommandHandle
}

func NewCommandHandlerMap() *CommandHandlerMap {
	c := CommandHandlerMap{}
	c.handlers = make(map[reflect.Type]CommandHandle, 0)

	return &c
}

func (c *CommandHandlerMap) Get(t reflect.Type) (CommandHandle, error) {
	if handler, exists := c.handlers[t]; exists {
		return handler, nil
	}

	return nil, fmt.Errorf("handler for {%v} not found", t)
}

func (c *CommandHandlerMap) AppendHandlers(commandHandlers ...CommandHandler) {
	for _, ch := range commandHandlers {
		for k, h := range ch.GetHandlers() {
			c.handlers[k] = h
		}
	}
}

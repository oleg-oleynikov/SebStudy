package infrastructure

import (
	"fmt"
	"reflect"
)

// TODO: Изменить на логику когда можешь убрать функционал удалив файл, ну или закоментив содержимое

type CommandHandlerModule interface {
	RegisterCommands(cmdHandlerMap *CommandHandlerMap)
}

type CommandHandlerMap struct {
	handlers map[reflect.Type]func(Command, CommandMetadata) error
}

func NewCommandHandlerMap() CommandHandlerMap {
	c := CommandHandlerMap{}
	c.handlers = make(map[reflect.Type]func(Command, CommandMetadata) error, 0)

	return c
}

func (c *CommandHandlerMap) Get(t reflect.Type) (func(Command, CommandMetadata) error, error) {
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

// func (c *CommandHandlerMap) RegisterCommand(valueType reflect.Type, f func(Command, CommandMetadata) error) {
// 	c.handlers[valueType] = f
// }

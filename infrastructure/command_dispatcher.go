package infrastructure

import (
	"fmt"
)

// var (
// 	instance *Dispatcher
// 	once     sync.Once
// )

type Dispatcher struct {
	commandHandlerMap CommandHandlerMap
}

func (d *Dispatcher) Dispatch(command interface{}, metadata CommandMetadata) error {
	handler, err := d.commandHandlerMap.Get(GetType(command))
	if err != nil {
		return fmt.Errorf("no handler registered")
	}

	return handler(command, metadata)
}

// func GetDispatcherInstance(commandHandlerMap CommandHandlerMap) *Dispatcher {
// 	once.Do(func() {
// 		instance = &Dispatcher{
// 			commandHandlerMap: commandHandlerMap,
// 		}
// 	})

// 	return instance
// }

func NewDispatcher(commandHandlerMap CommandHandlerMap) *Dispatcher {
	return &Dispatcher{
		commandHandlerMap: commandHandlerMap,
	}
}

package ports

import "SebStudy/infrastructure"

type CommandDispatcher interface {
	Dispatch(command interface{}, metadata infrastructure.CommandMetadata) error
}

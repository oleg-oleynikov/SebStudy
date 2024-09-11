package ports

import "SebStudy/internal/infrastructure"

type CommandDispatcher interface {
	Dispatch(command interface{}, metadata infrastructure.CommandMetadata) error
}

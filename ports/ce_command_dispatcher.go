package ports

import "SebStudy/infrastructure"

type CeCommandDispatcher interface {
	Dispatch(command interface{}, metadata infrastructure.CommandMetadata) error
}

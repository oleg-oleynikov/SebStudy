package ports

import "SebStudy/infrastructure"

type CeEventDispatcher interface {
	Dispatch(command interface{}, metadata infrastructure.CommandMetadata) error
}

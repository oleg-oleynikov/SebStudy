package ports

import "SebStudy/infrastructure"

type CeEventDispatcher interface {
	Dispatch(event interface{}, metadata infrastructure.EventMetadata) error
}

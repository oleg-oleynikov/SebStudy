package ports

import "SebStudy/infrastructure"

type CeEventHandler interface {
	Handle(event interface{}, metadata infrastructure.EventMetadata) error
}

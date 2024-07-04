package ports

import "SebStudy/eventsourcing"

type CeEventDispatcher interface {
	Dispatch(event interface{}, metadata eventsourcing.EventMetadata) error
}

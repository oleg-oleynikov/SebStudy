package infrastructure

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type CommandMetadata struct {
	CloudEvent cloudevents.Event
}

func NewCommandMetadataFromCloudEvent(cloudEvent cloudevents.Event) CommandMetadata {
	return CommandMetadata{
		CloudEvent: cloudEvent,
	}
}

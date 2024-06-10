package infrastructure

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type CommandMetadata struct {
	CloudEvent cloudevents.Event
	Context    *context.Context
}

func NewCommandMetadata(cloudevent cloudevents.Event, ctx *context.Context) CommandMetadata {
	return CommandMetadata{
		CloudEvent: cloudevent,
		Context:    ctx,
	}
}

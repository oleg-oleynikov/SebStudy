package commands

import (
	"context"
)

type CommandMetadata struct {
	// CloudEvent client.CloudEvents
	Context *context.Context
}

func NewCommandMetadata( /*cloudevent client.CloudEvents, */ ctx *context.Context) CommandMetadata {
	return CommandMetadata{
		// CloudEvent: cloudevent,
		Context: ctx,
	}
}

package infrastructure

import (
	v1 "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/options/grpc/protobuf/v1"
)

type CommandMetadata struct {
	CloudEvent *v1.CloudEvent
}

func NewCommandMetadataFromCloudEvent(cloudEvent *v1.CloudEvent) CommandMetadata {
	return CommandMetadata{
		CloudEvent: cloudEvent,
	}
}

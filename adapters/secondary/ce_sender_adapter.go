package secondary

import (
	"SebStudy/adapters/util"
	"SebStudy/infrastructure"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/options/grpc/protobuf/v1"
)

type CeSenderAdapter struct {
	Client *infrastructure.CloudeventsServiceClient

	CeMapper *util.CeMapper
}

func NewCeSenderAdapter(client *infrastructure.CloudeventsServiceClient, ceMapper *util.CeMapper) *CeSenderAdapter {
	return &CeSenderAdapter{
		Client:   client,
		CeMapper: ceMapper,
	}
}

func (c *CeSenderAdapter) SendEvent(e interface{}, eventType, source string) error {
	cloudEvent, err := c.newCloudEvent(e, eventType, source)

	if err != nil {
		return status.Errorf(codes.InvalidArgument, "failed to pack cloudevent for send event: %v", err)
	}

	if err := c.Client.Publish(cloudEvent); err != nil {
		return status.Errorf(codes.Internal, "failed to send cloudevent: %v", err)
	}

	return status.Errorf(codes.OK, "OK")
}

func (c *CeSenderAdapter) newCloudEvent(data interface{}, eventType, source string) (*v1.CloudEvent, error) {
	cloudEvent, err := c.CeMapper.MapToCloudEvent(data, eventType, source)
	if err != nil {
		return cloudEvent, err
	}

	return cloudEvent, err
}

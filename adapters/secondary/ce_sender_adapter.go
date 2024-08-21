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

	CloudeventMapper *util.CloudeventMapper
}

func NewCeSenderAdapter(client *infrastructure.CloudeventsServiceClient, cloudeventMapper *util.CloudeventMapper) *CeSenderAdapter {
	return &CeSenderAdapter{
		Client:           client,
		CloudeventMapper: cloudeventMapper,
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

	return nil
}

func (c *CeSenderAdapter) newCloudEvent(data interface{}, eventType, source string) (*v1.CloudEvent, error) {
	mapToCloudevent, err := c.CloudeventMapper.GetEventToCloudevent(infrastructure.GetValueType(data))

	if err != nil {
		return nil, err
	}

	return mapToCloudevent(eventType, source, data)
}

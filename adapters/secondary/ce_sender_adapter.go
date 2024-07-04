package secondary

import (
	"SebStudy/adapters/util"
	"context"
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"
)

type CeSenderAdapter struct {
	Client   cloudevents.Client
	Context  context.Context
	CeMapper *util.CeMapper
}

func NewCeSenderAdapter(targetUrl string, ceMapper *util.CeMapper) *CeSenderAdapter {
	client, context := newCloudEventsClient(targetUrl)
	return &CeSenderAdapter{
		Client:   client,
		Context:  context,
		CeMapper: ceMapper,
	}
}

func (c *CeSenderAdapter) SendEvent(e interface{}, eventType, source string) error {
	cloudEvent, err := c.newCloudEvent(e, eventType, source)

	if err != nil {
		return err
	}
	result := c.Client.Send(c.Context, cloudEvent)
	if cloudevents.IsUndelivered(result) {
		return fmt.Errorf("failed to send cloud event: %v", result)
	} else {
		var httpResult *cehttp.Result
		if cloudevents.ResultAs(result, &httpResult) {
			log.Printf("Sent with status code %d", httpResult.StatusCode)
		} else {
			return fmt.Errorf("send did not return an HTTP response: %s", result)
		}
	}

	// return nil

	return nil
}

func (c *CeSenderAdapter) newCloudEvent(data interface{}, eventType, source string) (cloudevents.Event, error) {
	cloudEvent, err := c.CeMapper.MapToCloudEvent(data, eventType, source)
	if err != nil {
		return cloudEvent, err
	}

	return cloudEvent, err
}

func newCloudEventsClient(targetUrl string) (cloudevents.Client, context.Context) {
	p, _ := cloudevents.NewHTTP()

	c, err := cloudevents.NewClient(p, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	// if err != nil {
	// 	log.Fatalf("failed to create protocol: %s", err.Error())
	// }

	// c, err := cloudevents.NewClient(p, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	// if err != nil {
	// 	log.Fatalf("failed to create client, %v", err)
	// }
	ctx := cloudevents.ContextWithTarget(context.Background(), targetUrl)

	return c, ctx
}

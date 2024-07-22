package primary

import (
	"SebStudy/adapters/util"
	"SebStudy/infrastructure"
	"SebStudy/ports"
	"context"
	"log"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/gorilla/handlers"
)

type CloudEventsAdapter struct {
	CommandDispatcher ports.CeCommandDispatcher
	EventDispatcher   ports.CeEventHandler
	// Client            cloudevents.Client
	CeMapper *util.CeMapper
}

func NewCloudEventsAdapter(d ports.CeCommandDispatcher, e ports.CeEventHandler, ceMapper *util.CeMapper) *CloudEventsAdapter {
	return &CloudEventsAdapter{
		CommandDispatcher: d,
		EventDispatcher:   e,

		// Client:   newCloudEventsClient(),
		CeMapper: ceMapper,
	}
}

// func newCloudEventsClient() cloudevents.Client {
// 	p, err := cloudevents.NewHTTP()
// 	if err != nil {
// 		log.Fatalf("failed to create protocol: %s", err.Error())
// 	}

// 	ce, err := cloudevents.NewClient(p)
// 	if err != nil {
// 		log.Fatalf("failed to create http client, %v", err)
// 	}

// 	return ce
// }

func (c *CloudEventsAdapter) ceHandler(w http.ResponseWriter, r *http.Request) {
	event, err := cloudevents.NewEventFromHTTPRequest(r)
	if err != nil {
		log.Printf("failed to parse CloudEvent from request: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ceEvent := *event

	if _, err := c.CeMapper.GetEventType(event.Type()); err != nil {
		log.Printf("unknown event type: %s\n", err)
		http.Error(w, "Unknown event type:"+err.Error(), http.StatusBadRequest)
		return
	}

	mappedEvent, err := c.CeMapper.MapToEvent(context.Background(), ceEvent)

	if err != nil {
		log.Printf("failed to map cloudevent: %v", err)
		http.Error(w, "failed to map cloudevent: "+err.Error(), http.StatusBadRequest)
		return
	}

	if c.CeMapper.IsCommand(event.Type()) {
		err = c.CommandDispatcher.Dispatch(mappedEvent, infrastructure.NewCommandMetadataFromCloudEvent(ceEvent))
		if err != nil {
			if _, ok := err.(cloudevents.Result); ok {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			log.Printf("failed to dispatch command: %v", err)
			http.Error(w, "failed to dispatch command: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else if c.CeMapper.IsEvent(event.Type()) {
		err := c.EventDispatcher.Handle(mappedEvent, *infrastructure.NewEventMetadataFromCloudEvent(ceEvent))
		if err != nil {
			if _, ok := err.(cloudevents.Result); ok {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			log.Printf("failed to handle event: %v", err)
			http.Error(w, "failed to handle event: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	log.Println("Unknown request")
	http.Error(w, "Unknown request", http.StatusBadRequest)
}

func (c *CloudEventsAdapter) Run() {

	go func() {
		ceHandler := handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"}),
			handlers.AllowedHeaders([]string{"Content-Type"}),
		)(http.HandlerFunc(c.ceHandler))

		http.Handle("/", ceHandler)

		log.Println("Start server on localhost:8080")
		http.ListenAndServe(":8080", nil)
		// log.Fatalf("failed to start receiver: %s", c.Client.StartReceiver(context.Background(), c.receive))
	}()
}

// func (c *CloudEventsAdapter) receiveCe(ctx context.Context, event cloudevents.Event) cloudevents.Result {
// 	// event.Data()
// 	// log.Println("Пришло что то нахуй")
// 	// log.Println(event)

// 	if _, err := c.CeMapper.GetEventType(event.Type()); err != nil {
// 		log.Printf("unknown event type: %s\n", err)
// 		return cloudevents.NewHTTPResult(http.StatusBadRequest, "Unknown event type: %s", err)
// 	}

// 	mappedEvent, err := c.CeMapper.MapToEvent(ctx, event)

// 	if err != nil {
// 		log.Printf("failed to map cloudevent: %v", err)
// 		return cloudevents.NewHTTPResult(http.StatusBadRequest, "failed to map cloudevent: %v", err)
// 	}

// 	if c.CeMapper.IsCommand(event.Type()) {
// 		err = c.CommandDispatcher.Dispatch(mappedEvent, infrastructure.NewCommandMetadataFromCloudEvent(event))
// 		if err != nil {
// 			if _, ok := err.(cloudevents.Result); ok {
// 				return err
// 			}

// 			log.Printf("failed to dispatch command: %v", err)
// 			return cloudevents.NewHTTPResult(500, "failed to dispatch command: %v", err)
// 		}
// 	} else if c.CeMapper.IsEvent(event.Type()) {
// 		err := c.EventDispatcher.Handle(mappedEvent, *infrastructure.NewEventMetadataFromCloudEvent(event))
// 		if err != nil {
// 			if _, ok := err.(cloudevents.Result); ok {
// 				return err
// 			}

// 			log.Printf("failed to handle event: %v", err)
// 			return cloudevents.NewHTTPResult(http.StatusInternalServerError, "failed to handle event: %v", err)
// 		}
// 	}

// 	log.Println("Unknown request")
// 	return cloudevents.ResultNACK
// }

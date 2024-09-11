package adapters

import (
	"SebStudy/logger"
	"SebStudy/pb"
	"fmt"
	"sync"
)

type EventProcessor struct {
	eventChan         chan *pb.CloudEvent
	wg                sync.WaitGroup
	log               logger.Logger
	subscriberManager *SubscriberManager
}

func NewEventProcessor(chanSize int, log logger.Logger, subscriberManager *SubscriberManager) *EventProcessor {
	return &EventProcessor{
		eventChan:         make(chan *pb.CloudEvent, chanSize),
		log:               log,
		subscriberManager: subscriberManager,
	}
}

func (ep *EventProcessor) SubmitEvent(event *pb.CloudEvent) error {
	select {
	case ep.eventChan <- event:
		return nil
	default:
		ep.log.Debugf("EventChan is full, dropping event: %v", event)
		return fmt.Errorf("EventChan is full, event dropped")
	}
}

func (ep *EventProcessor) StartProcessing() {
	for {
		select {
		case event := <-ep.eventChan:
			ep.wg.Add(1)
			go func(event *pb.CloudEvent) {
				defer ep.wg.Done()
				ep.subscriberManager.BroadcastEvent(event)
			}(event)
		default:
			return
		}
	}
}

package adapters

import (
	"SebStudy/logger"
	"SebStudy/pb"
	"sync"
)

type SubscriberManager struct {
	sync.RWMutex
	subscribers map[pb.CloudEventService_SubscribeServer]chan *pb.CloudEvent
	log         logger.Logger
}

func NewSubcriberManager(log logger.Logger) *SubscriberManager {
	return &SubscriberManager{
		subscribers: make(map[pb.CloudEventService_SubscribeServer]chan *pb.CloudEvent),
		log:         log,
	}
}

func (sm *SubscriberManager) AddSubscriber(stream pb.CloudEventService_SubscribeServer) {
	sm.Lock()
	defer sm.Unlock()

	eventChan := make(chan *pb.CloudEvent, 100)
	sm.subscribers[stream] = eventChan

	go func() {
		for {
			event, ok := <-eventChan
			if !ok {
				return
			}
			if err := stream.Send(event); err != nil {
				sm.log.Printf("Failed to send event to subscriber: %v", err)
				sm.RemoveSubscriber(stream)
				return
			}
		}
	}()
}

func (sm *SubscriberManager) RemoveSubscriber(subscriber pb.CloudEventService_SubscribeServer) {
	sm.Lock()
	defer sm.Unlock()

	if eventChan, ok := sm.subscribers[subscriber]; ok {
		subscriber.Context().Done()
		close(eventChan)
		delete(sm.subscribers, subscriber)
	}
}

func (sm *SubscriberManager) BroadcastEvent(event *pb.CloudEvent) {
	sm.RLock()
	defer sm.RUnlock()

	for _, eventChan := range sm.subscribers {
		select {
		case eventChan <- event:
		default:
			sm.log.Printf("Subscriber queue is full, dropping event")
		}
	}
}

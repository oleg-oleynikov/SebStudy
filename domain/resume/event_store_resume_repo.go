package resume

import (
	"SebStudy/domain/resume/values"
	"SebStudy/infrastructure"
)

type EventStoreResumeRepo struct {
	eventStore *infrastructure.EventStore
}

func NewEventStoreResumeRepo(eventStore *infrastructure.EventStore) *EventStoreResumeRepo {
	return &EventStoreResumeRepo{
		eventStore: eventStore,
	}
}

func (es *EventStoreResumeRepo) Get(resumeId *values.ResumeId) (*Resume, error) {
	return nil, nil

}

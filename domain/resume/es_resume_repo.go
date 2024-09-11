package resume

import (
	"SebStudy/domain/resume/values"
	"SebStudy/infrastructure"
	"SebStudy/infrastructure/eventsourcing"
)

type EsResumeRepository struct {
	aggregateStore eventsourcing.AggregateStore
}

func NewEsResumeRepository(aggregateStore eventsourcing.AggregateStore) *EsResumeRepository {
	return &EsResumeRepository{
		aggregateStore: aggregateStore,
	}
}

func (es *EsResumeRepository) Get(resumeId *values.ResumeId) (*Resume, error) {
	resume := NewResume()
	err := es.aggregateStore.Load(resumeId.Value, resume)

	if err != nil {
		return nil, err
	}

	return resume, nil
}

func (es *EsResumeRepository) Save(resume *Resume, m infrastructure.CommandMetadata) error {
	return es.aggregateStore.Save(resume, m)
}

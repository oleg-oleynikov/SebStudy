package resume

import (
	"SebStudy/domain/resume/values"
	"SebStudy/infrastructure"
)

type EsResumeRepo struct {
	aggregateStore infrastructure.AggregateStore
}

func NewEsResumeRepo(aggregateStore infrastructure.AggregateStore) *EsResumeRepo {
	return &EsResumeRepo{
		aggregateStore: aggregateStore,
	}
}

func (es *EsResumeRepo) Get(resumeId *values.ResumeId) (*Resume, error) {
	resume := NewResume()
	err := es.aggregateStore.Load(resumeId.Value, resume)

	if err != nil {
		return nil, err
	}

	return resume, nil
}

func (es *EsResumeRepo) Save(resume *Resume, m infrastructure.CommandMetadata) error {
	return es.aggregateStore.Save(resume, m)
}

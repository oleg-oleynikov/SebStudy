package resume

import (
	"resume-server/internal/domain/resume/models"
	"resume-server/internal/domain/resume/values"
	"resume-server/internal/infrastructure"
	"resume-server/internal/infrastructure/eventsourcing"
)

type EsResumeRepository struct {
	aggregateStore eventsourcing.AggregateStore
}

func NewEsResumeRepository(aggregateStore eventsourcing.AggregateStore) *EsResumeRepository {
	return &EsResumeRepository{
		aggregateStore: aggregateStore,
	}
}

func (es *EsResumeRepository) Get(resumeId *values.ResumeId) (*models.Resume, error) {
	resume := models.NewResume()
	err := es.aggregateStore.Load(resumeId.Id, resume)

	if err != nil {
		return nil, err
	}

	return resume, nil
}

func (es *EsResumeRepository) Save(resume *models.Resume, m infrastructure.CommandMetadata) error {
	return es.aggregateStore.Save(resume, m)
}

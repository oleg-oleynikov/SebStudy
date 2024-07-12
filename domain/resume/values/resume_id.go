package values

import (
	"fmt"
)

type ResumeId struct {
	// resumeId uuid.UUID
	Value string
}

func NewResumeId(resumeId string) *ResumeId {
	// id := uuid.New()
	return &ResumeId{
		Value: resumeId,
	}
}

func (rID *ResumeId) ToString() string {
	return fmt.Sprintf("%v", rID)
}

package values

import (
	"fmt"
)

type ResumeId struct {
	// resumeId uuid.UUID
	Value int
}

func NewResumeId(resumeId int) *ResumeId {
	// id := uuid.New()
	return &ResumeId{
		Value: resumeId,
	}
}

func (rID *ResumeId) ToString() string {
	return fmt.Sprintf("%v", rID)
}

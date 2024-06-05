package values

import (
	"fmt"

	"github.com/google/uuid"
)

type ResumeId struct {
	resumeId uuid.UUID
	// 	resumeId int
}

func NewResumeId(resumeId int) *ResumeId {
	id := uuid.New()
	return &ResumeId{
		resumeId: resumeId,
	}
}

func (rID *ResumeId) ToString() string {
	return fmt.Sprintf("%v", rID)
}

package values

import (
	"fmt"
	//"github.com/google/uuid"
)

type ResumeId struct {
	// resumeId uuid.UUID
	resumeId int
}

func NewResumeId(rID int) *ResumeId {
	return &ResumeId{
		resumeId: rID,
	}
}

func (rID *ResumeId) ToString() string {
	return fmt.Sprintf("%v", rID)
}

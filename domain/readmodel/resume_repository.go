package readmodel

import "SebStudy/domain/resume"

type ResumeRepository interface {
	GetResume(resumeId string) (*resume.Resume, error)
}

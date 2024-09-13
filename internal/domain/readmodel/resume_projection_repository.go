package readmodel

import (
	"SebStudy/internal/domain/resume"
)

type ResumeProjectionRepository interface {
	AddResume(resume.Resume) error
	UpdateResume(resume.Resume) error
}

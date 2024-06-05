package commands

import (
	"SebStudy/domain/resume"
)

type SendResume struct {
	resume resume.Resume
}

func NewSendResume(resume resume.Resume) *SendResume {
	return &SendResume{
		resume: resume,
	}
}

package commands

import (
	resume "SebStudy/domain"
	"time"
)

type SendResume struct {
	resume    resume.Resume
	timestamp time.Time
}

func NewSendResume(resume resume.Resume, timestamp time.Time) *SendResume {
	return &SendResume{
		resume:    resume,
		timestamp: timestamp,
	}
}

package commands

import domain "SebStudy/domain/resume"

type SendResume struct {
	domain.Resume
}

func NewSendResume(resume domain.Resume) SendResume {
	return SendResume{
		Resume: resume,
	}
}

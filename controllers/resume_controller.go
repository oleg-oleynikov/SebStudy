package controllers

import (
	"SebStudy/infrastructure/handlers/commands"
)

type ResumeController struct {
	commandDispatcher *commands.Dispatcher
}

func NewResumeController(d *commands.Dispatcher) *ResumeController {
	return &ResumeController{
		commandDispatcher: d,
	}
}

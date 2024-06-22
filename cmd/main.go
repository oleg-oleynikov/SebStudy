package main

import (
	"SebStudy/adapters/primary"
	"SebStudy/domain/resume"
	"SebStudy/infrastructure"
)

const (
	port = 8080
)

func main() {
	handlers := resume.NewHandlers(nil)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap(handlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)

	ceMapper := primary.NewCeMapper()

	s := primary.NewCloudEventsAdapter(dispatcher, ceMapper, port)

	s.Run()
}

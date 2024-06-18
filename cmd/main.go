package main

import (
	"SebStudy/adapters/primary"
	"SebStudy/domain/resume"
	"SebStudy/infrastructure"
)

func main() {
	// disp :=
	handlers := resume.NewHandlers(nil)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap(handlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)

	ceMapper := primary.NewCeMapper()

	s := primary.NewCloudEventsAdapter(dispatcher, ceMapper, 8080)

	s.Run()
	// u := uuid.New()
	// fmt.Println(u)
}

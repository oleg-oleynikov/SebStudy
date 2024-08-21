package mapping

import (
	"SebStudy/adapters/util"
	"SebStudy/domain/resume/events"
	"reflect"
)

func RegisterResumeTypes(m *util.CloudeventMapper) {
	m.MapEvent(
		reflect.TypeOf(events.ResumeCreated{}),
		"resume.created",
		util.EVENT,
		toResumeCreated,
		toCloudeventResumeCreated,
	)

	m.MapCommand(
		"resume.create",
		util.COMMAND,
		toCreateResumeCommand,
	)
}

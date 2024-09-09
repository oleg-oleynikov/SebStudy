package mapping

import (
	"SebStudy/util"
)

func RegisterResumeTypes(m *util.CloudEventCommandAdapter) {
	m.MapCommand(
		"resume.create",
		toCreateResume,
	)
}

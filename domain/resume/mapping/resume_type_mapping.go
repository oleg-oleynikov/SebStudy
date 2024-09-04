package mapping

import (
	"SebStudy/adapters/util"
)

func RegisterResumeTypes(m *util.CloudeventMapper) {

	m.MapCommand(
		"resume.create",
		toCreateResume,
	)

}

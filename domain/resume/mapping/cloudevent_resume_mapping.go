package mapping

import (
	"SebStudy/util"
)

func RegisterCloudeventResumeTypes(m *util.CloudEventCommandAdapter) {

	m.MapCloudevent(
		"type.googleapis.com/resume.ResumeCreated",
		toCreateResume,
	)

}

package mapping

import (
	"SebStudy/internal/infrastructure/util"
)

func RegisterCloudeventResumeTypes(m *util.CloudEventCommandAdapter) {

	m.MapCloudevent(
		"type.googleapis.com/resume.ResumeCreated",
		toCreateResume,
	)

}

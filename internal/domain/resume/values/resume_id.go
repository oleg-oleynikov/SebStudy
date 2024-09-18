package values

type ResumeId struct {
	Id string
}

func NewResumeId(resumeId string) ResumeId {
	return ResumeId{
		Id: resumeId,
	}
}

func (rID *ResumeId) String() string {
	return rID.Id
}

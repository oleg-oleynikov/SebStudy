package models

type ResumeProjection struct {
	Id            string
	Education     string   `bson:"education"`
	AboutMe       string   `bson:"aboutMe,omitempty"`
	BornDate      uint64   `bson:"bornDate"`
	Skills        []string `bson:"skills,omitempty"`
	Direction     string   `bson:"direction"`
	AboutProjects string   `bson:"aboutProjects,omitempty"`
	Portfolio     string   `bson:"portfolio,omitempty"`
}

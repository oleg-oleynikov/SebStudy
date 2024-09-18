package models

import "time"

type ResumeProjection struct {
	Id            string    `bson:"_id,omitempty"`
	Education     string    `bson:"education,omitempty"`
	AboutMe       string    `bson:"aboutMe,omitempty"`
	BirthDate     time.Time `bson:"birthDate,omitempty"`
	Skills        []string  `bson:"skills,omitempty"`
	Direction     string    `bson:"direction,omitempty"`
	AboutProjects string    `bson:"aboutProjects,omitempty"`
	Portfolio     string    `bson:"portfolio,omitempty"`
	UserId        string    `bson:"userId,omitempty"`
}

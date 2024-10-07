package models

import (
	"resume-server/pb"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

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

func ResumeProjectionToProto(rp *ResumeProjection) *pb.Resume {
	skills := []*pb.Skill{}
	for _, skill := range rp.Skills {
		skills = append(skills, &pb.Skill{Skill: skill})
	}

	return &pb.Resume{
		ResumeId:      rp.Id,
		Education:     rp.Education,
		AboutMe:       rp.AboutMe,
		Skills:        skills,
		BirthDate:     timestamppb.New(rp.BirthDate),
		Direction:     rp.Direction,
		AboutProjects: rp.AboutProjects,
		Portfolio:     rp.Portfolio,
	}
}

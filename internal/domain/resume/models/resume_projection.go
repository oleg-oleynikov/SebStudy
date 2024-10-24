package models

import (
	"resume-server/pb"
)

type ResumeProjection struct {
	Id            string   `bson:"_id,omitempty"`
	AboutMe       string   `bson:"aboutMe,omitempty"`
	Skills        []string `bson:"skills,omitempty"`
	Direction     string   `bson:"direction,omitempty"`
	AboutProjects string   `bson:"aboutProjects,omitempty"`
	Portfolio     string   `bson:"portfolio,omitempty"`
	UserId        string   `bson:"userId,omitempty"`
}

func ResumeProjectionToProto(rp *ResumeProjection) *pb.Resume {
	skills := []*pb.Skill{}
	for _, skill := range rp.Skills {
		skills = append(skills, &pb.Skill{Skill: skill})
	}

	return &pb.Resume{
		ResumeId:      rp.Id,
		AboutMe:       rp.AboutMe,
		Skills:        skills,
		Direction:     rp.Direction,
		AboutProjects: rp.AboutProjects,
		Portfolio:     rp.Portfolio,
	}
}

package grpc

import (
	"SebStudy/internal/domain/resume/commands"
	"SebStudy/internal/domain/resume/service"
	"SebStudy/internal/domain/resume/values"
	"SebStudy/internal/infrastructure"
	"SebStudy/logger"
	"SebStudy/pb"
	"context"
	"time"

	"github.com/gofrs/uuid"
)

type resumeGrpcService struct {
	pb.UnimplementedResumeServiceServer

	rs  *service.ResumeService
	log logger.Logger
}

func NewResumeGrpcService(log logger.Logger, rs *service.ResumeService) *resumeGrpcService {
	return &resumeGrpcService{log: log, rs: rs}
}

func (s *resumeGrpcService) CreateResume(ctx context.Context, req *pb.CreateResumeReq) (*pb.CreateResumeRes, error) {
	aggregateId, _ := uuid.NewV7()

	education, err := values.NewEducation(req.GetEducation())
	if err != nil {
		s.log.Debugf("Failed to create education: %v", err)
		return nil, err
	}

	aboutMe, err := values.NewAboutMe(req.GetAboutMe())
	if err != nil {
		s.log.Debugf("Failed to create aboutMe: %v", err)
		return nil, err
	}

	skills, _ := values.NewSkills()
	for i := 0; i < len(req.Skills); i++ {
		data := req.Skills[i]
		skill, err := values.NewSkill(data.Skill)
		if err != nil {
			s.log.Debugf("Failed to create skill: %v", err)
			return nil, err
		}
		skills.AppendSkills(*skill)
	}

	timeBirth, err := time.Parse("2006-01-02", req.GetBirthDate())
	if err != nil {
		s.log.Debugf("Failed to create timeBirth: %v", err)
		return nil, err
	}

	birthDate, err := values.NewBirthDate(timeBirth)
	if err != nil {
		s.log.Debugf("Failed to create: %v", err)
		return nil, err
	}

	direction, err := values.NewDirection(req.GetDirection())
	if err != nil {
		s.log.Debugf("Failed to create direction: %v", err)
		return nil, err
	}

	aboutProjects, err := values.NewAboutProjects(req.GetAboutProjects())
	if err != nil {
		s.log.Debugf("Failed to create aboutProjects: %v", err)
		return nil, err
	}

	portfolio, err := values.NewPortfolio(req.GetPortfolio())
	if err != nil {
		s.log.Debugf("Failed to create portfolio: %v", err)
		return nil, err
	}

	command := commands.NewCreateResume(aggregateId.String(), *education, *aboutMe, *skills, *birthDate, *direction, *aboutProjects, *portfolio)

	md := infrastructure.NewCommandMetadata(aggregateId.String())

	if err := s.rs.Commands.CreateResume.Handle(ctx, command, md); err != nil {
		return nil, err
	}

	return &pb.CreateResumeRes{
		AggregateId: aggregateId.String(),
	}, nil
}

func (s *resumeGrpcService) ChangeResume(ctx context.Context, req *pb.ChangeResumeReq) (*pb.ChangeResumeRes, error) {
	md := infrastructure.NewCommandMetadata(req.GetResumeId())

	education, err := values.NewEducation(req.GetEducation())
	if err != nil {
		s.log.Debugf("Failed to create education: %v", err)
		return nil, err
	}

	aboutMe, err := values.NewAboutMe(req.GetAboutMe())
	if err != nil {
		s.log.Debugf("Failed to create aboutMe: %v", err)
		return nil, err
	}

	skills, _ := values.NewSkills()
	for i := 0; i < len(req.Skills); i++ {
		data := req.Skills[i]
		skill, err := values.NewSkill(data.Skill)
		if err != nil {
			s.log.Debugf("Failed to create skill: %v", err)
			return nil, err
		}
		skills.AppendSkills(*skill)
	}

	birthDate := &values.BirthDate{}
	if req.GetBirthDate() != "" {
		timeBirth, err := time.Parse("2006-01-02", req.GetBirthDate())
		if err != nil {
			s.log.Debugf("Failed to create timeBirth: %v", err)
			return nil, err
		}

		birthDate, err = values.NewBirthDate(timeBirth)
		if err != nil {
			s.log.Debugf("Failed to create: %v", err)
			return nil, err
		}
	}

	direction, err := values.NewDirection(req.GetDirection())
	if err != nil {
		s.log.Debugf("Failed to create direction: %v", err)
		return nil, err
	}

	aboutProjects, err := values.NewAboutProjects(req.GetAboutProjects())
	if err != nil {
		s.log.Debugf("Failed to create aboutProjects: %v", err)
		return nil, err
	}

	portfolio, err := values.NewPortfolio(req.GetPortfolio())
	if err != nil {
		s.log.Debugf("Failed to create portfolio: %v", err)
		return nil, err
	}

	command := commands.NewChangeResume(req.GetResumeId(), *education, *aboutMe, *skills, *birthDate, *direction, *aboutProjects, *portfolio)

	if err := s.rs.Commands.ChangeResume.Handle(ctx, command, md); err != nil {
		return nil, err
	}

	return &pb.ChangeResumeRes{}, nil
}

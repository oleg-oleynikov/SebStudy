package grpc

import (
	"context"
	"fmt"
	"log"
	"resume-server/internal/domain/resume/commands"
	"resume-server/internal/domain/resume/models"
	"resume-server/internal/domain/resume/queries"
	"resume-server/internal/domain/resume/service"
	"resume-server/internal/domain/resume/values"
	"resume-server/internal/infrastructure"
	"resume-server/logger"
	"resume-server/pb"
	"resume-server/pkg/interceptors"
	"time"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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
	AccountId, ok := ctx.Value(interceptors.AccountIdKey).(string)
	if !ok {
		s.log.Debugf("CreateResume. Failed to get AccountId from context")
		return nil, fmt.Errorf("failed to get AccountId")
	}

	query := queries.NewResumeExistsByAccountId(AccountId)
	exists, err := s.rs.Queries.ResumeExistsByAccountId.Handle(ctx, query)
	if err != nil {
		s.log.Debugf("CreateResume. Failed to check resume contains in read model: %v", err)
		return nil, status.Error(codes.Internal, "failed to check resume existence")
	}

	if exists {
		return nil, status.Error(codes.AlreadyExists, "resume already exists")
	}

	aggregateId, _ := uuid.NewV7()

	education, err := values.NewEducation(req.GetEducation())
	if err != nil {
		s.log.Debugf("Failed to create education: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	aboutMe, err := values.NewAboutMe(req.GetAboutMe())
	if err != nil {
		s.log.Debugf("Failed to create aboutMe: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	skills, _ := values.NewSkills()
	for i := 0; i < len(req.Skills); i++ {
		data := req.Skills[i]
		skill, err := values.NewSkill(data.Skill)
		if err != nil {
			s.log.Debugf("Failed to create skill: %v", err)
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		skills.AppendSkills(*skill)
	}

	timeBirth, err := time.Parse("2006-01-02", req.GetBirthDate())
	if err != nil {
		s.log.Debugf("Failed to create timeBirth: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	birthDate, err := values.NewBirthDate(timeBirth)
	if err != nil {
		s.log.Debugf("Failed to create: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	direction, err := values.NewDirection(req.GetDirection())
	if err != nil {
		s.log.Debugf("Failed to create direction: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	aboutProjects, err := values.NewAboutProjects(req.GetAboutProjects())
	if err != nil {
		s.log.Debugf("Failed to create aboutProjects: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	portfolio, err := values.NewPortfolio(req.GetPortfolio())
	if err != nil {
		s.log.Debugf("Failed to create portfolio: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	command := commands.NewCreateResume(aggregateId.String(), *education, *aboutMe, *skills, *birthDate, *direction, *aboutProjects, *portfolio)

	md := infrastructure.NewCommandMetadata(aggregateId.String(), AccountId)

	if err := s.rs.Commands.CreateResume.Handle(ctx, command, md); err != nil {
		return nil, err
	}

	return &pb.CreateResumeRes{
		AggregateId: aggregateId.String(),
	}, nil
}

func (s *resumeGrpcService) ChangeResume(ctx context.Context, req *pb.ChangeResumeReq) (*pb.ChangeResumeRes, error) {
	accountId, ok := ctx.Value(interceptors.AccountIdKey).(string)
	if !ok {
		s.log.Debugf("ChangeResume. Failed to get AccountId from context")
		return nil, fmt.Errorf("failed to get AccountId")
	}

	md := infrastructure.NewCommandMetadata(req.GetResumeId(), accountId)

	education, err := values.NewEducation(req.GetEducation())
	if err != nil {
		s.log.Debugf("Failed to create education: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	aboutMe, err := values.NewAboutMe(req.GetAboutMe())
	if err != nil {
		s.log.Debugf("Failed to create aboutMe: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	skills, _ := values.NewSkills()
	for i := 0; i < len(req.Skills); i++ {
		data := req.Skills[i]
		skill, err := values.NewSkill(data.Skill)
		if err != nil {
			s.log.Debugf("Failed to create skill: %v", err)
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		skills.AppendSkills(*skill)
	}

	birthDate := &values.BirthDate{}
	if req.GetBirthDate() != "" {
		timeBirth, err := time.Parse("2006-01-02", req.GetBirthDate())
		log.Println("timeBirth")
		if err != nil {
			s.log.Debugf("Failed to create timeBirth: %v", err)
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		birthDate, err = values.NewBirthDate(timeBirth)
		if err != nil {
			s.log.Debugf("Failed to create: %v", err)
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	direction, err := values.NewDirection(req.GetDirection())
	if err != nil {
		s.log.Debugf("Failed to create direction: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	aboutProjects, err := values.NewAboutProjects(req.GetAboutProjects())
	if err != nil {
		s.log.Debugf("Failed to create aboutProjects: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	portfolio, err := values.NewPortfolio(req.GetPortfolio())
	if err != nil {
		s.log.Debugf("Failed to create portfolio: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	command := commands.NewChangeResume(req.GetResumeId(), *education, *aboutMe, *skills, *birthDate, *direction, *aboutProjects, *portfolio)

	if err := s.rs.Commands.ChangeResume.Handle(ctx, command, md); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.ChangeResumeRes{}, nil
}

func (s *resumeGrpcService) GetResumeByAccountId(ctx context.Context, empty *emptypb.Empty) (*pb.GetResumeByAccountIdRes, error) {
	accountId, ok := ctx.Value(interceptors.AccountIdKey).(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "GetResumeByAccountId. Missing token")
	}

	query := queries.NewGetResumeByAccountIdQuery(accountId)
	rp, err := s.rs.Queries.GetResumeByAccountId.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return &pb.GetResumeByAccountIdRes{
		Resume: models.ResumeProjectionToProto(rp),
	}, nil
}

func (s *resumeGrpcService) GetResumeById(ctx context.Context, req *pb.GetResumeByIdReq) (*pb.GetResumeByIdRes, error) {
	query := queries.NewGetResumeByIdQuery(req.GetResumeId())
	rp, err := s.rs.Queries.GetResumeById.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return &pb.GetResumeByIdRes{
		Resume: models.ResumeProjectionToProto(rp),
	}, nil
}

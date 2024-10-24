package grpc

import (
	"context"
	"fmt"
	"resume-server/internal/domain/resume/commands"
	"resume-server/internal/domain/resume/models"
	"resume-server/internal/domain/resume/queries"
	"resume-server/internal/domain/resume/service"
	"resume-server/internal/domain/resume/values"
	"resume-server/internal/infrastructure"
	"resume-server/logger"
	"resume-server/pb"
	"resume-server/pkg/interceptors"

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

func (s *resumeGrpcService) CreateResume(ctx context.Context, req *pb.CreateResumeReq) (*pb.ResumeResponse, error) {
	AccountId, ok := ctx.Value(interceptors.AccountIdKey).(string)
	if !ok {
		s.log.Debugf("CreateResume. Failed to get AccountId from context")
		return nil, fmt.Errorf("failed to get AccountId")
	}
	s.log.Debugf("AccountId: %s", AccountId)

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

	command := commands.NewCreateResume(aggregateId.String(), *aboutMe, *skills, *direction, *aboutProjects, *portfolio)

	md := infrastructure.NewCommandMetadata(aggregateId.String(), AccountId)

	if err := s.rs.Commands.CreateResume.Handle(ctx, command, md); err != nil {
		return nil, err
	}

	resume := &pb.Resume{
		ResumeId:      aggregateId.String(),
		AboutMe:       aboutMe.GetAboutMe(),
		Skills:        skills.ToProto(),
		Direction:     direction.GetDirection(),
		AboutProjects: aboutProjects.GetAboutProjects(),
		Portfolio:     portfolio.GetPortfolio(),
	}

	return &pb.ResumeResponse{
		Resume: resume,
	}, nil
}

func (s *resumeGrpcService) ChangeResume(ctx context.Context, req *pb.ChangeResumeReq) (*pb.ResumeResponse, error) {
	accountId, ok := ctx.Value(interceptors.AccountIdKey).(string)
	if !ok {
		s.log.Debugf("ChangeResume. Failed to get AccountId from context")
		return nil, fmt.Errorf("failed to get AccountId")
	}

	md := infrastructure.NewCommandMetadata(req.GetResumeId(), accountId)

	query := queries.NewGetResumeByAccountIdQuery(accountId)
	rp, err := s.rs.Queries.GetResumeByAccountId.Handle(ctx, query)
	if err != nil {
		s.log.Debugf("(GetResumeByAccountIdQuery) Error by query: %v", err)
		return nil, fmt.Errorf("user and resume dont match")
	} else {
		if rp.Id != req.GetResumeId() {
			s.log.Debugf("ChangeResume. Resume does not belong to the user")
			return nil, fmt.Errorf("resume does not belong to the user")
		}
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

	command := commands.NewChangeResume(req.GetResumeId(), *aboutMe, *skills, *direction, *aboutProjects, *portfolio)

	if err := s.rs.Commands.ChangeResume.Handle(ctx, command, md); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resume := &pb.Resume{
		ResumeId:      req.GetResumeId(),
		AboutMe:       aboutMe.GetAboutMe(),
		Skills:        skills.ToProto(),
		Direction:     direction.GetDirection(),
		AboutProjects: aboutProjects.GetAboutProjects(),
		Portfolio:     portfolio.GetPortfolio(),
	}

	return &pb.ResumeResponse{
		Resume: resume,
	}, nil
}

func (s *resumeGrpcService) GetResumeByAccountId(ctx context.Context, empty *emptypb.Empty) (*pb.ResumeResponse, error) {
	accountId, ok := ctx.Value(interceptors.AccountIdKey).(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "GetResumeByAccountId. Missing token")
	}

	query := queries.NewGetResumeByAccountIdQuery(accountId)
	rp, err := s.rs.Queries.GetResumeByAccountId.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return &pb.ResumeResponse{
		Resume: models.ResumeProjectionToProto(rp),
	}, nil
}

func (s *resumeGrpcService) GetResumeById(ctx context.Context, req *pb.GetResumeByIdReq) (*pb.ResumeResponse, error) {
	query := queries.NewGetResumeByIdQuery(req.GetResumeId())
	rp, err := s.rs.Queries.GetResumeById.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return &pb.ResumeResponse{
		Resume: models.ResumeProjectionToProto(rp),
	}, nil
}

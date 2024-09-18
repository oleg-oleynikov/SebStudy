package grpc

import (
	"SebStudy/logger"
	"SebStudy/pb"
	"context"
)

type resumeGrpcService struct {
	pb.UnimplementedResumeServiceServer
	log logger.Logger
}

func NewResumeGrpcService(log logger.Logger) *resumeGrpcService {
	return &resumeGrpcService{log: log}
}

func (s *resumeGrpcService) CreateResume(context.Context, *pb.CreateResumeReq) (*pb.CreateResumeRes, error) {
	// aggregateID := uuid.NewV7().String()
	return nil, nil
}

func (s *resumeGrpcService) ChangeResume(context.Context, *pb.ChangeResumeReq) (*pb.ChangeResumeRes, error) {
	return nil, nil
}

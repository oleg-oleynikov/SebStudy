package server

import (
	"net"

	delivery "resume-server/internal/domain/resume/delivery/grpc"
	"resume-server/pb"
	"resume-server/pkg/interceptors"

	"google.golang.org/grpc"
)

func (s *server) NewResumeGrpcServer() (func() error, *grpc.Server, error) {
	l, err := net.Listen("tcp", s.cfg.GRPC.Port)
	if err != nil {
		return nil, nil, err
	}

	interceptorManager := interceptors.NewInterceptorManager(s.log, s.cfg)

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptorManager.AuthInterceptor))

	resumeService := delivery.NewResumeGrpcService(s.log, s.rs)
	pb.RegisterResumeServiceServer(grpcServer, resumeService)

	go func() {
		s.log.Infof("gRPC server is listening on port {%s}", s.cfg.GRPC.Port)
		s.log.Error(grpcServer.Serve(l))
	}()

	return l.Close, grpcServer, nil
}

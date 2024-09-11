package server

import (
	"SebStudy/adapters"
	"SebStudy/pb"
	"net"

	"google.golang.org/grpc"
)

func (s *server) NewResumeGrpcServer() (func() error, *grpc.Server, error) {
	l, err := net.Listen("tcp", s.cfg.ServerPort)
	if err != nil {
		return nil, nil, err
	}

	grpcServer := grpc.NewServer()
	cloudeventService := adapters.NewCloudEventService(s.log, s.cmdDispatcher, s.cmdAdapter)
	pb.RegisterCloudEventServiceServer(grpcServer, cloudeventService)

	go func() {
		s.log.Infof("gRPC server is listening on port {%s}", s.cfg.GRPC.Port)
		s.log.Error(grpcServer.Serve(l))
	}()

	return l.Close, grpcServer, nil
}

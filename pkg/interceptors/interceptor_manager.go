package interceptors

import (
	"SebStudy/config"
	"SebStudy/logger"
	"SebStudy/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InterceptorManager struct {
	log logger.Logger
	cfg *config.Config

	authClient pb.SSOServerServiceClient
}

func NewInterceptorManager(log logger.Logger, cfg *config.Config) *InterceptorManager {
	conn, err := grpc.NewClient(cfg.AuthServerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
		return nil
	}

	client := pb.NewSSOServerServiceClient(conn)

	return &InterceptorManager{
		log: log,
		cfg: cfg,

		authClient: client,
	}
}

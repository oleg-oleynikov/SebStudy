package interceptors

import (
	"SebStudy/pb"
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const AccountIdKey = contextKey("accountId")

func (m *InterceptorManager) AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	token := md["authorization"]
	if len(token) == 0 {
		m.log.Debugf("AuthInterceptor. Missing token")
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}

	accessToken := strings.TrimPrefix(token[0], "Bearer ")
	if accessToken == "" {
		m.log.Debugf("AuthInterceptor. Missing token")
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}

	accountId, err := m.callVerifyToken(ctx, accessToken)
	if err != nil {
		m.log.Debugf("AuthInterceptor. Failed to call verify token: %v")
		return nil, status.Error(codes.Internal, "AuthInterceptor. Failed to call verify token")
	}

	newCtx := context.WithValue(ctx, AccountIdKey, accountId)

	return handler(newCtx, req)
}

func (m *InterceptorManager) callVerifyToken(ctx context.Context, accessToken string) (string, error) {
	in := &pb.Token{
		Token: accessToken,
	}

	res, err := m.authClient.VerifyToken(ctx, in)
	if err != nil {
		return "", err
	}

	if !res.GetStatus() {
		return "", fmt.Errorf("invalid token")
	}

	return res.GetAccountId(), nil
}

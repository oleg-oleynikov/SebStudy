package interceptors

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	claims, err := m.verifyToken(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to get claims: %v", err)
	}

	accountId, ok := claims["subjectId"].(string)
	if !ok || accountId == "" {
		m.log.Debugf("AuthInterceptor. Failed to get subjectId")
		return nil, status.Error(codes.Unauthenticated, "missing subject")
	}

	newCtx := context.WithValue(ctx, AccountIdKey, accountId)

	return handler(newCtx, req)
}

func (m *InterceptorManager) verifyToken(token string) (jwt.MapClaims, error) {
	publicKey, err := m.getPublicKey()
	if err != nil {
		m.log.Debugf("Failed to read public key: %v", err)
		return nil, fmt.Errorf("failed to read public key")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		exp, ok := claims["exp"].(float64)
		if !ok {
			return nil, errors.New("invalid token expiration time")
		}
		if int64(exp) < time.Now().Unix() {
			return nil, errors.New("token is expired")
		}

		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (m *InterceptorManager) getPublicKey() (*rsa.PublicKey, error) {
	keyPath := m.cfg.PublicKeyPath
	readKey, _ := os.ReadFile(keyPath)
	block, _ := pem.Decode(readKey)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		m.log.Debugf("failed to parse public key", err)
		return nil, err
	}
	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		m.log.Fatalf("not an RSA public key")
	}

	return publicKey, nil
}

// func (m *InterceptorManager) callVerifyToken(ctx context.Context, accessToken string) (string, error) {
// 	in := &pb.Token{
// 		Token: accessToken,
// 	}

// 	res, err := m.authClient.VerifyToken(ctx, in)
// 	if err != nil {
// 		m.log.Debugf("(AuthInterceptor) Error: %v", err)
// 		return "", err
// 	}

// 	if !res.GetStatus() {
// 		m.log.Debugf("(AuthInterceptor) invalid token")
// 		return "", fmt.Errorf("invalid token")
// 	}

// 	return res.GetAccountId(), nil
// }

package primary

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type corsCheck func(md metadata.MD) error

type grpcCors struct {
	corsChecks []corsCheck
}

func (c *grpcCors) addCorsCheck(handler corsCheck) {
	c.corsChecks = append(c.corsChecks, handler)
}

func NewGrpcCorsBuilder() *grpcCors {
	return &grpcCors{}
}

func (c *grpcCors) WithAllowedOrigins(origins ...string) *grpcCors {
	corsCheck := func(md metadata.MD) error {

		origin := md.Get("Origin")
		if len(origin) == 0 {
			return status.Errorf(codes.InvalidArgument, "missing 'Origin' header")
		}

		if origins != nil && !containsString(origin[0], origins) {
			return status.Errorf(codes.PermissionDenied, "origin not allowed: %s", origin[0])
		}

		return nil
	}

	c.addCorsCheck(corsCheck)
	return c
}

func (c *grpcCors) WithAllowedMethods(methods ...string) *grpcCors {
	return nil
}

func (c *grpcCors) WithAllowedHeaders(headers ...string) *grpcCors {
	return nil
}

func (c *grpcCors) Build() func(ctx context.Context, method string) error {
	return func(ctx context.Context, method string) error {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Errorf(codes.Internal, "failed to get metadata")
		}

		for _, h := range c.corsChecks {
			if err := h(md); err != nil {
				return err
			}
		}

		return nil
	}
}

func containsString(needle string, haystack []string) bool {
	for _, s := range haystack {
		if s == needle || s == "*" {
			return true
		}
	}
	return false
}

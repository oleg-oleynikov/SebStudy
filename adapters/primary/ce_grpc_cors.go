package primary

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type corsCheck func(md metadata.MD) error
type CorsHandler func(ctx context.Context, method string) error

type CorsGrpc struct {
	corsChecks []corsCheck
}

func (c *CorsGrpc) addCorsCheck(handler corsCheck) {
	c.corsChecks = append(c.corsChecks, handler)
}

func NewCorsGrpcBuilder() *CorsGrpc {
	return &CorsGrpc{}
}

func (c *CorsGrpc) WithAllowedOrigins(origins ...string) *CorsGrpc {
	corsCheckAllowedOrigin := func(md metadata.MD) error {

		origin := md.Get("Origin")
		if len(origin) == 0 {
			return status.Errorf(codes.InvalidArgument, "missing 'Origin' header")
		}

		if origins != nil && !containsString(origin[0], origins) {
			return status.Errorf(codes.PermissionDenied, "origin not allowed: %s", origin[0])
		}

		return nil
	}

	c.addCorsCheck(corsCheckAllowedOrigin)
	return c
}

// func (c *CorsGrpc) WithAllowedMethods(methods ...string) *CorsGrpc {
// 	if len(methods) == 0 {
// 		log.Printf("allowed methods is empty")
// 		return c
// 	}
// 	corsCheckAllowedMethods := func(md metadata.MD) error {
// 		log.Println(md.Get("method")[0])
// 		return nil
// 	}

// 	c.addCorsCheck(corsCheckAllowedMethods)
// 	return c
// }

func (c *CorsGrpc) WithAllowedHeaders(headers ...string) *CorsGrpc {
	if len(headers) == 0 {
		log.Printf("allowed headers is empty")
		return c
	}

	return c
}

func (c *CorsGrpc) BuildHandler() CorsHandler {
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

func CorsToServerOptions(corsHandler CorsHandler) []grpc.ServerOption {
	unaryCorsInterceptor := grpc.UnaryInterceptor(BuildCorsUnaryInterceptor(corsHandler))
	streamCorsInterceptor := grpc.StreamInterceptor(BuildCorsStreamInterceptor(corsHandler))
	return []grpc.ServerOption{unaryCorsInterceptor, streamCorsInterceptor}
}

func BuildCorsUnaryInterceptor(cors CorsHandler) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := cors(ctx, info.FullMethod); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func BuildCorsStreamInterceptor(cors func(ctx context.Context, method string) error) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if err := cors(ss.Context(), info.FullMethod); err != nil {
			return err
		}
		return handler(srv, ss)
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

package interceptor

import (
	"context"
	"fmt"
	"strings"

	"github.com/Tel3scop/chat-server/internal/connector/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const authPrefix = "Bearer "

type validator interface {
	Validate() error
}

// CheckAuth интерсептор, который позволяет валидировать запрос, если присутствует метод Validate
func CheckAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, fmt.Errorf("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, fmt.Errorf("invalid authorization header format")
	}
	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	if info == nil {
		return nil, fmt.Errorf("can not get path")
	}

	err := auth.CheckAuth(accessToken, info.FullMethod)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

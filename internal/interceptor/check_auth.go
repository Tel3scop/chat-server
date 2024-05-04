package interceptor

import (
	"context"
	"fmt"

	"github.com/Tel3scop/chat-server/internal/connector/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Client contains client connection with authentication service.
type Client struct {
	Client *auth.Client
}

// CheckAuth интерсептор, который позволяет валидировать запрос, если присутствует метод Validate
func (c *Client) CheckAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}

	if info == nil {
		return nil, fmt.Errorf("can not get path")
	}

	err := c.Client.Check(metadata.NewOutgoingContext(ctx, md), info.FullMethod)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

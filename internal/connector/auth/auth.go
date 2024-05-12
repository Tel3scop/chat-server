package auth

import (
	"context"
	"fmt"

	"github.com/Tel3scop/auth/pkg/access_v1"
	"github.com/Tel3scop/chat-server/internal/connector"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client экземпляр
type Client struct {
	client access_v1.AccessV1Client
}

var _ connector.AuthClient = (*Client)(nil)

// New создает новый экземпляр клиента
func New(host string, port int64) (*Client, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial GRPC client: %v", err)
	}

	return &Client{
		client: access_v1.NewAccessV1Client(conn),
	}, nil
}

// Check calls authentication service method for authorization.
func (c *Client) Check(ctx context.Context, endpoint string) error {
	_, err := c.client.Check(ctx, &access_v1.CheckRequest{
		EndpointAddress: endpoint,
	})
	return err
}

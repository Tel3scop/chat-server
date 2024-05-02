package auth

import (
	"context"
	"fmt"

	"github.com/Tel3scop/chat-server/pkg/access_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var Connection *grpc.ClientConn

func New(host string, port int64) error {
	var err error
	Connection, err = grpc.Dial(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to dial GRPC client: %v", err)
	}
	return nil
}

// CheckAuth вызов метода Check из пакета auth-service
func CheckAuth(token, endpoint string) error {
	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": "Bearer " + token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	//пока его закинул в pkg, как заапрувишь auth, просто через go get добавлю пакет
	cl := access_v1.NewAccessV1Client(Connection)

	_, err := cl.Check(ctx, &access_v1.CheckRequest{
		EndpointAddress: endpoint,
	})
	if err != nil {
		return err
	}

	fmt.Println("Access granted")
	return nil
}

package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/Tel3scop/chat-server/pkg/access_v1"

	"github.com/Tel3scop/chat-server/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// CheckAuth вызов метода Check из пакета auth-service
func CheckAuth(cfg *config.Config, token, endpoint string) error {
	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": "Bearer " + token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.AuthService.Host, cfg.AuthService.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial GRPC client: %v", err)
	}

	defer conn.Close()

	//пока его закинул в pkg, как заапрувишь auth, просто через go get добавлю пакет
	cl := access_v1.NewAccessV1Client(conn)

	_, err = cl.Check(ctx, &access_v1.CheckRequest{
		EndpointAddress: endpoint,
	})
	if err != nil {
		return err
	}

	fmt.Println("Access granted")
	return nil
}

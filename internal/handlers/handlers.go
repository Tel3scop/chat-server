package handlers

import (
	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run запуск всех хендлеров
func Run() *grpc.Server {
	s := grpc.NewServer()
	reflection.Register(s)
	chatAPI.RegisterChatV1Server(s, &chatServer{})

	return s
}

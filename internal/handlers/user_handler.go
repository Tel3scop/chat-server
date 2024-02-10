package handlers

import (
	"context"
	"log"

	"github.com/Tel3scop/chat-server/internal/services/chat_service"
	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type chatServer struct {
	chatAPI.UnimplementedChatV1Server
}

func (s *chatServer) Create(ctx context.Context, req *chatAPI.CreateRequest) (*chatAPI.CreateResponse, error) {
	log.Printf("Creating data: %+v", req)
	createdChatID := chat_service.Create(ctx, req)
	return &chatAPI.CreateResponse{Id: createdChatID}, nil
}

func (s *chatServer) Delete(ctx context.Context, req *chatAPI.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Deleting data: %+v", req)
	chat_service.DeleteByID(ctx, req.GetId())
	return &emptypb.Empty{}, nil
}

func (s *chatServer) SendMessage(ctx context.Context, req *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Sending message: %+v", req)
	chat_service.SendMessage(ctx, req)
	return &emptypb.Empty{}, nil
}

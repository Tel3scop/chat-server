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

	createdChatID, err := chat_service.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &chatAPI.CreateResponse{Id: createdChatID}, nil
}

func (s *chatServer) Delete(ctx context.Context, req *chatAPI.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Deleting data: %+v", req)

	err := chat_service.DeleteByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *chatServer) SendMessage(ctx context.Context, req *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Sending message: %+v", req)

	err := chat_service.SendMessage(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

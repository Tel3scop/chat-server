package chat

import (
	"context"

	"github.com/Tel3scop/chat-server/internal/api/converter"
	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
)

// Create создание нового чата
func (i *Implementation) Create(ctx context.Context, req *chatAPI.CreateRequest) (*chatAPI.CreateResponse, error) {
	createdChatID, err := i.chatService.Create(ctx, converter.ToChatModelFromRequest(req))
	if err != nil {
		return nil, err
	}

	return &chatAPI.CreateResponse{Id: createdChatID}, err
}

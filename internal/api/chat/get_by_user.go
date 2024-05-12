package chat

import (
	"context"

	"github.com/Tel3scop/chat-server/internal/api/converter"
	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
)

// GetChats получить все чаты пользователя
func (i *Implementation) GetChats(ctx context.Context, req *chatAPI.GetChatsRequest) (*chatAPI.GetChatsResponse, error) {

	chats, err := i.chatService.GetChatsByUsername(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	return converter.ToChatsResponseFromModel(chats), err
}

package chat

import (
	"context"

	"github.com/Tel3scop/chat-server/internal/api/converter"
	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
)

// GetMessages метод получения сообщений
func (i *Implementation) GetMessages(ctx context.Context, req *chatAPI.GetMessagesRequest) (*chatAPI.GetMessagesResponse, error) {
	messages, err := i.chatService.GetMessagesByChatID(ctx, req.GetChatId(), req.GetCount())
	if err != nil {
		return nil, err
	}

	return converter.ToMessagesResponseFromModel(messages), err
}

package chat

import (
	"context"

	"github.com/Tel3scop/chat-server/internal/api/converter"
	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage метод отправки сообщения
func (i *Implementation) SendMessage(ctx context.Context, req *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.messageService.SendMessage(ctx, req.ChatId, converter.ToMessageModelFromRequest(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}

package chat

import (
	"context"
	"log"

	"github.com/Tel3scop/chat-server/internal/api/converter"
	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage метод отправки сообщения
func (i *Implementation) SendMessage(ctx context.Context, req *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Sending message: %+v", req)
	requestModel := converter.ToMessageModelFromRequest(req)
	err := i.messageService.SendMessage(ctx, req.ChatId, requestModel)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}

package chat

import (
	"context"

	"github.com/Tel3scop/chat-server/internal/api/converter"
	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage метод отправки сообщения
func (i *Implementation) SendMessage(ctx context.Context, req *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	i.mxChannel.RLock()
	chatChan, ok := i.channels[req.GetChatId()]
	i.mxChannel.RUnlock()
	if !ok {
		err := i.chatService.CheckChatByUsernameAndID(ctx, req.From, req.GetChatId())
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "chat not found")
		}
		i.channels[req.GetChatId()] = make(chan *chatAPI.Message, 100)
	}

	err := i.chatService.SendMessage(ctx, req.ChatId, converter.ToMessageModelFromRequest(req))
	if err != nil {
		return nil, err
	}

	go func() {
		chatChan <- &chatAPI.Message{
			From:      req.From,
			Text:      req.Text,
			CreatedAt: req.Timestamp,
		}
	}()

	return &emptypb.Empty{}, err
}

package chat

import (
	"context"
	"log"

	"github.com/Tel3scop/chat-server/internal/model"
)

// SendMessage отправка сообщения в чат
func (s *serv) SendMessage(ctx context.Context, chatID int64, message model.Message) error {

	err := s.chatRepository.SendMessage(ctx, chatID, message)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

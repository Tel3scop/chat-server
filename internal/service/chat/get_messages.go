package chat

import (
	"context"
	"log"

	"github.com/Tel3scop/chat-server/internal/model"
)

// GetMessagesByChatID получить сообщения в чате
func (s *serv) GetMessagesByChatID(ctx context.Context, chatID, count int64) ([]model.Message, error) {
	messages, err := s.chatRepository.GetMessagesByChatID(ctx, chatID, count)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return messages, nil
}

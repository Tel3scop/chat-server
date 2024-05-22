package chat

import (
	"context"
	"log"

	"github.com/Tel3scop/chat-server/internal/model"
)

// GetChatsByUsername получить все доступные пользователю чаты
func (s *serv) GetChatsByUsername(ctx context.Context, username string) ([]model.Chat, error) {
	chats, err := s.chatRepository.GetChatsByUsername(ctx, username)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return chats, nil
}

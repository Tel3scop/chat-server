package chat

import (
	"context"

	"github.com/Tel3scop/chat-server/internal/model"
)

// Create создание нового чата. В транзакции создаем чат и првиязываем к нему переданных пользователей
func (s *serv) Create(ctx context.Context, chat model.Chat) (int64, error) {
	var chatID int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		chatID, errTx = s.chatRepository.Create(ctx, chat)
		if errTx != nil {
			return errTx
		}
		errTx = s.chatRepository.LinkUsers(ctx, chatID, chat.Usernames, chat.CreatedAt)
		if errTx != nil {
			return errTx
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

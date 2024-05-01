package chat

import (
	"context"
	"log"

	"github.com/Tel3scop/chat-server/internal/client/db"
	"github.com/Tel3scop/chat-server/internal/model"
	"github.com/Tel3scop/chat-server/internal/repository"
	"github.com/Tel3scop/chat-server/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

// NewService функция возвращает новый сервис чата
func NewService(
	chatRepository repository.ChatRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}

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

// DeleteByID удаление чата по ID, удаляет через ON CASCADE все связанные сущности: пользователей и сообщения
func (s *serv) DeleteByID(ctx context.Context, id int64) error {
	err := s.chatRepository.DeleteByID(ctx, id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

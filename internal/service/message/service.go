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
	messageRepository repository.MessageRepository
	txManager         db.TxManager
}

// NewService функция возвращает новый сервис сообщений
func NewService(
	messageRepository repository.MessageRepository,
	txManager db.TxManager,
) service.MessageService {
	return &serv{
		messageRepository: messageRepository,
		txManager:         txManager,
	}
}

// SendMessage отправка сообщения в чат
func (s *serv) SendMessage(ctx context.Context, chatID int64, message model.Message) error {
	err := s.messageRepository.SendMessage(ctx, chatID, message)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

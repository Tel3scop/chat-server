package chat

import (
	"github.com/Tel3scop/chat-server/internal/client/db"
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

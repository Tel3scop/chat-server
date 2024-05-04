package service

import (
	"context"

	"github.com/Tel3scop/chat-server/internal/model"
)

// ChatService интерфейс для использования в сервисе
type ChatService interface {
	Create(ctx context.Context, chatData model.Chat) (int64, error)
	DeleteByID(ctx context.Context, id int64) error
}

// MessageService интерфейс для использования в сервисе
type MessageService interface {
	SendMessage(ctx context.Context, chatID int64, message model.Message) error
}

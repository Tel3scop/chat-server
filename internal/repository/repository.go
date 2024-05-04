package repository

import (
	"context"
	"time"

	"github.com/Tel3scop/chat-server/internal/model"
)

// ChatRepository интерфейс репозитория пользователей
type ChatRepository interface {
	Create(ctx context.Context, chatData model.Chat) (int64, error)
	DeleteByID(ctx context.Context, id int64) error
	LinkUsers(ctx context.Context, chatID int64, usernames []string, createdAt time.Time) error
}

// MessageRepository интерфейс репозитория истории изменения
type MessageRepository interface {
	SendMessage(ctx context.Context, chatID int64, message model.Message) error
}

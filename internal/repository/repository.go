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
	SendMessage(ctx context.Context, chatID int64, message model.Message) error
	CheckChatByUsernameAndID(ctx context.Context, username string, ID int64) error
	GetChatsByUsername(ctx context.Context, username string) ([]model.Chat, error)
	GetMessagesByChatID(ctx context.Context, chatID, count int64) ([]model.Message, error)
}

package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Tel3scop/chat-server/internal/client/db"
	"github.com/Tel3scop/chat-server/internal/model"
	"github.com/Tel3scop/chat-server/internal/repository"
)

const (
	tableName = "messages"

	columnID        = "id"
	columnText      = "text"
	columnUsername  = "username"
	columnChatID    = "chat_id"
	columnCreatedAt = "created_at"
)

type repo struct {
	db db.Client
}

// NewRepository создание репозитория для пользователей
func NewRepository(db db.Client) repository.MessageRepository {
	return &repo{db: db}
}

// SendMessage отправка нового сообщения в чат
func (r *repo) SendMessage(ctx context.Context, chatID int64, message model.Message) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(columnChatID, columnText, columnUsername, columnCreatedAt).
		Values(chatID, message.Text, message.From, message.Timestamp)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "message_repository.SendMessage",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

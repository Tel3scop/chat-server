package chat

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/Tel3scop/chat-server/internal/client/db"
	"github.com/Tel3scop/chat-server/internal/model"
	"github.com/Tel3scop/chat-server/internal/repository"
)

const (
	tableName = "chats"

	columnID        = "id"
	columnName      = "name"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepository создание репозитория для пользователей
func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

// Create Метод создания нового чата.
func (r *repo) Create(ctx context.Context, dto model.Chat) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(columnName, columnCreatedAt, columnUpdatedAt).
		Values(dto.Name, dto.CreatedAt, dto.UpdatedAt).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// DeleteByID удаление чата.
func (r *repo) DeleteByID(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).PlaceholderFormat(sq.Dollar).Where(sq.Eq{columnID: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("error building delete query: %s", err)
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}
	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("error executing delete query: %s", err)
	}

	if result.RowsAffected() > 0 {
		return nil
	}

	return fmt.Errorf("can not delete record")
}

// LinkUsers Метод привязывает юзеров к конкретному чату.
func (r *repo) LinkUsers(ctx context.Context, chatID int64, usernames []string, createdAt time.Time) error {
	builder := sq.Insert("chat_user").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "username", "created_at")
	for _, username := range usernames {
		builder = builder.Values(chatID, username, createdAt)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.LinkUsers",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

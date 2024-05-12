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

// таблица chats
const (
	tableName = "chats"

	columnID        = "id"
	columnName      = "name"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
)

// таблица messages
const (
	tableMessages = "messages"

	columnMessageID        = "id"
	columnMessageText      = "text"
	columnMessageUsername  = "username"
	columnMessageChatID    = "chat_id"
	columnMessageCreatedAt = "created_at"
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

// CheckChatByUsernameAndID проверка существования чата по ID и пользователю.
func (r *repo) CheckChatByUsernameAndID(ctx context.Context, username string, ID int64) error {
	builder := sq.Select(columnID).From(tableName).PlaceholderFormat(sq.Dollar).
		Join("chat_user on chats.ID = chat_user.chat_id").
		Where(sq.Eq{columnID: ID}).
		Where(sq.Eq{"username": username}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("error building delete query: %s", err)
	}

	q := db.Query{
		Name:     "chat_repository.CheckChatByUsernameAndID",
		QueryRaw: query,
	}
	var chatID int64
	err = r.db.DB().ScanOneContext(ctx, &chatID, q, args...)
	if err != nil {
		return fmt.Errorf("error executing query: %s", err)
	}

	return nil
}

// SendMessage отправка нового сообщения в чат
func (r *repo) SendMessage(ctx context.Context, chatID int64, message model.Message) error {
	builder := sq.Insert(tableMessages).
		PlaceholderFormat(sq.Dollar).
		Columns(columnMessageChatID, columnMessageText, columnMessageUsername, columnMessageCreatedAt).
		Values(chatID, message.Text, message.From, message.Timestamp)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.SendMessage",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetChatsByUsername получить все доступные пользователю чаты.
func (r *repo) GetChatsByUsername(ctx context.Context, username string) ([]model.Chat, error) {
	builder := sq.Select(columnID, columnName, "array_agg(chat_user.username) as usernames").From(tableName).PlaceholderFormat(sq.Dollar).
		Join("chat_user on chats.ID = chat_user.chat_id").
		Join("chat_user filter on chats.ID = filter.chat_id").
		Where(sq.Eq{"filter.username": username}).
		GroupBy(columnID)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building delete query: %s", err)
	}

	q := db.Query{
		Name:     "chat_repository.GetChatsByUsername",
		QueryRaw: query,
	}
	var chats []model.Chat
	err = r.db.DB().ScanAllContext(ctx, &chats, q, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %s", err)
	}

	return chats, nil
}

// GetMessagesByChatID получить count последних сообщений по chatID.
func (r *repo) GetMessagesByChatID(ctx context.Context, chatID, count int64) ([]model.Message, error) {
	builder := sq.Select(columnMessageUsername, columnMessageText, columnMessageCreatedAt).From(tableMessages).PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnMessageChatID: chatID}).
		Limit(uint64(count)).
		OrderBy(columnMessageCreatedAt + " DESC")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building delete query: %s", err)
	}

	q := db.Query{
		Name:     "chat_repository.GetMessagesByChatID",
		QueryRaw: query,
	}
	var messages []model.Message
	err = r.db.DB().ScanAllContext(ctx, &messages, q, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %s", err)
	}

	return messages, nil
}

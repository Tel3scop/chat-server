package chat_service

import (
	"context"
	"log"

	"github.com/Tel3scop/chat-server/internal/entities"
	"github.com/Tel3scop/chat-server/internal/storages/chat_storage"
	"github.com/Tel3scop/chat-server/pkg/chat_v1"
)

// Create создание нового чата.
func Create(ctx context.Context, request *chat_v1.CreateRequest) int64 {
	createdChat, err := chat_storage.Create(ctx, entities.Chat{Usernames: request.Usernames})
	if err != nil {
		return 0
	}
	return createdChat.ID
}

// DeleteByID удаление чата по ID
func DeleteByID(ctx context.Context, id int64) {
	err := chat_storage.DeleteByID(ctx, id)
	if err != nil {
		log.Println(err.Error())
	}
}

// SendMessage отправка сообщения в чат
func SendMessage(ctx context.Context, request *chat_v1.SendMessageRequest) {
	message := entities.Message{
		From:      request.From,
		Text:      request.Text,
		Timestamp: request.Timestamp.AsTime(),
	}
	err := chat_storage.SendMessage(ctx, request.ChatId, message)
	if err != nil {
		log.Println(err.Error())
	}
}

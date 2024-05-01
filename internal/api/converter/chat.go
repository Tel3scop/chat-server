package converter

import (
	"time"

	"github.com/Tel3scop/chat-server/internal/model"
	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
)

// ToMessageModelFromRequest функция получения ответа из модели
func ToMessageModelFromRequest(request *chatAPI.SendMessageRequest) model.Message {
	return model.Message{
		From:      request.From,
		Text:      request.Text,
		Timestamp: time.Now(),
	}
}

// ToChatModelFromRequest функция получения модели пользователя из запроса
func ToChatModelFromRequest(request *chatAPI.CreateRequest) model.Chat {
	now := time.Now()
	return model.Chat{
		Usernames: request.Usernames,
		Name:      request.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

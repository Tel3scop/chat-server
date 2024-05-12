package converter

import (
	"time"

	"github.com/Tel3scop/chat-server/internal/model"
	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
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

// ToChatsResponseFromModel функция получения ответа с чатами из модели
func ToChatsResponseFromModel(data []model.Chat) *chatAPI.GetChatsResponse {
	response := chatAPI.GetChatsResponse{Chats: make([]*chatAPI.Chat, 0, len(data))}
	for _, chat := range data {
		response.Chats = append(response.Chats, &chatAPI.Chat{
			Id:        chat.ID,
			Name:      chat.Name,
			Usernames: chat.Usernames,
		})
	}
	return &response
}

// ToMessagesResponseFromModel функция получения ответа с сообщениями из модели
func ToMessagesResponseFromModel(data []model.Message) *chatAPI.GetMessagesResponse {
	response := chatAPI.GetMessagesResponse{Messages: make([]*chatAPI.Message, 0, len(data))}
	for _, message := range data {
		response.Messages = append(response.Messages, &chatAPI.Message{
			From:      message.From,
			Text:      message.Text,
			CreatedAt: timestamppb.New(message.Timestamp),
		})
	}
	return &response
}

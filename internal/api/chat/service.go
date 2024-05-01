package chat

import (
	"github.com/Tel3scop/chat-server/internal/service"
	"github.com/Tel3scop/chat-server/pkg/chat_v1"
)

// Implementation структура для работы с хэндерами авторизации
type Implementation struct {
	chat_v1.UnimplementedChatV1Server
	chatService    service.ChatService
	messageService service.MessageService
}

// NewImplementation новый экземпляр структуры Implementation
func NewImplementation(chatService service.ChatService, messageService service.MessageService) *Implementation {
	return &Implementation{
		chatService:    chatService,
		messageService: messageService,
	}
}

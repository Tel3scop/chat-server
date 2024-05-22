package chat

import (
	"sync"

	"github.com/Tel3scop/chat-server/internal/service"
	"github.com/Tel3scop/chat-server/pkg/chat_v1"
)

// Implementation структура для работы с хэндерами чата
type Implementation struct {
	chat_v1.UnimplementedChatV1Server
	chatService service.ChatService
	chats       map[int64]*Chat
	mxChat      sync.RWMutex

	channels  map[int64]chan *chat_v1.Message
	mxChannel sync.RWMutex
}

// Chat структура чата со стримами
type Chat struct {
	streams map[string]chat_v1.ChatV1_ConnectChatServer
	m       sync.RWMutex
}

// NewImplementation новый экземпляр структуры Implementation
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
		chats:       make(map[int64]*Chat),
		channels:    make(map[int64]chan *chat_v1.Message),
	}
}

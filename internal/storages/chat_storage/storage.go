package chat_storage

import (
	"context"
	"fmt"
	"sync"

	"github.com/Tel3scop/chat-server/internal/entities"
)

// SyncMap Эмуляция БД с сиквенсом
type SyncMap struct {
	elems    map[int64]entities.Chat
	m        sync.RWMutex
	sequence int64
}

var chats = &SyncMap{
	elems: make(map[int64]entities.Chat),
}

// Create Метод создания нового чата.
func Create(ctx context.Context, chatData entities.Chat) (entities.Chat, error) {
	_ = ctx
	chats.m.Lock()
	defer chats.m.Unlock()
	chats.sequence++
	chatData.ID = chats.sequence
	chats.elems[chats.sequence] = chatData
	return chats.elems[chats.sequence], nil
}

// SendMessage отправка нового сообщения в чат
func SendMessage(ctx context.Context, chatID int64, message entities.Message) error {
	_ = ctx
	chats.m.Lock()
	defer chats.m.Unlock()

	chat, ok := chats.elems[chatID]
	if !ok {
		return fmt.Errorf("chat %d not found", chatID)
	}
	chat.Messages = append(chat.Messages, message)
	chats.elems[chatID] = chat

	return nil
}

// DeleteByID удаление чата.
func DeleteByID(ctx context.Context, id int64) error {
	_ = ctx
	chats.m.Lock()
	defer chats.m.Unlock()
	_, ok := chats.elems[id]
	if !ok {
		return fmt.Errorf("chat %d not found", id)
	}
	delete(chats.elems, id)

	return nil
}

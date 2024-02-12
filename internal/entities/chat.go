package entities

import "time"

// Chat Структура чата
type Chat struct {
	ID        int64     `json:"id"`
	Usernames []string  `json:"usernames"`
	Messages  []Message `json:"messages"`
}

// Message Структура сообщения
type Message struct {
	From      string    `json:"from"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

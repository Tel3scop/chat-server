package model

import "time"

// Chat Структура чата
type Chat struct {
	ID        int64     `json:"id" db:"id"`
	Usernames []string  `json:"usernames" db:"usernames"`
	Messages  []Message `json:"messages"`
	Name      string    `json:"name" db:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Message Структура сообщения
type Message struct {
	From      string    `json:"from" db:"username"`
	Text      string    `json:"text" db:"text"`
	Timestamp time.Time `json:"timestamp" db:"created_at"`
}

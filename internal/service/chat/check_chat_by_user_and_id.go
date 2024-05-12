package chat

import (
	"context"
	"log"
)

// CheckChatByUsernameAndID проверка чата по пользователю и ID чата
func (s *serv) CheckChatByUsernameAndID(ctx context.Context, username string, ID int64) error {
	err := s.chatRepository.CheckChatByUsernameAndID(ctx, username, ID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

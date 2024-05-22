package chat

import (
	"context"

	"github.com/Tel3scop/helpers/logger"
)

// CheckChatByUsernameAndID проверка чата по пользователю и ID чата
func (s *serv) CheckChatByUsernameAndID(ctx context.Context, username string, ID int64) error {
	err := s.chatRepository.CheckChatByUsernameAndID(ctx, username, ID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

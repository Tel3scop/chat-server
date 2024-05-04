package chat

import (
	"context"
	"log"
)

// DeleteByID удаление чата по ID, удаляет через ON CASCADE все связанные сущности: пользователей и сообщения
func (s *serv) DeleteByID(ctx context.Context, id int64) error {
	err := s.chatRepository.DeleteByID(ctx, id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

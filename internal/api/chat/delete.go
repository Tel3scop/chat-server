package chat

import (
	"context"
	"log"

	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete удаление чата и связанных сущностей
func (i *Implementation) Delete(ctx context.Context, req *chatAPI.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Deleting data: %+v", req)
	err := i.chatService.DeleteByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}

package chat

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	chatAPI "github.com/Tel3scop/chat-server/pkg/chat_v1"
)

// ConnectChat метод подключения к чату
func (i *Implementation) ConnectChat(req *chatAPI.ConnectChatRequest, stream chatAPI.ChatV1_ConnectChatServer) error {
	i.mxChannel.RLock()
	chatChan, ok := i.channels[req.GetChatId()]
	i.mxChannel.RUnlock()
	if !ok {
		ctx := stream.Context()
		err := i.chatService.CheckChatByUsernameAndID(ctx, req.Username, req.GetChatId())
		if err != nil {
			return status.Errorf(codes.NotFound, "chat not found")
		}
		i.channels[req.GetChatId()] = make(chan *chatAPI.Message, 100)
	}

	i.mxChat.Lock()
	if _, okChat := i.chats[req.GetChatId()]; !okChat {
		i.chats[req.GetChatId()] = &Chat{
			streams: make(map[string]chatAPI.ChatV1_ConnectChatServer),
		}
	}
	i.mxChat.Unlock()

	i.chats[req.GetChatId()].m.Lock()
	i.chats[req.GetChatId()].streams[req.GetUsername()] = stream
	i.chats[req.GetChatId()].m.Unlock()

	for {
		select {
		case msg, okCh := <-chatChan:
			if !okCh {
				return nil
			}

			for userReceiver, st := range i.chats[req.GetChatId()].streams {
				if msg.From == userReceiver {
					continue
				}

				if err := st.Send(msg); err != nil {
					return err
				}
			}

		case <-stream.Context().Done():
			i.chats[req.GetChatId()].m.Lock()
			delete(i.chats[req.GetChatId()].streams, req.GetUsername())
			i.chats[req.GetChatId()].m.Unlock()
			return nil
		}
	}
}

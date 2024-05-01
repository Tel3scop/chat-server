package app

import (
	"context"
	"log"

	"github.com/Tel3scop/chat-server/internal/api/chat"
	"github.com/Tel3scop/chat-server/internal/client/db"
	"github.com/Tel3scop/chat-server/internal/client/db/pg"
	"github.com/Tel3scop/chat-server/internal/client/db/transaction"
	"github.com/Tel3scop/chat-server/internal/closer"
	"github.com/Tel3scop/chat-server/internal/config"
	"github.com/Tel3scop/chat-server/internal/repository"
	chatRepo "github.com/Tel3scop/chat-server/internal/repository/chat"
	messageRepo "github.com/Tel3scop/chat-server/internal/repository/message"
	"github.com/Tel3scop/chat-server/internal/service"
	chatService "github.com/Tel3scop/chat-server/internal/service/chat"
	messageService "github.com/Tel3scop/chat-server/internal/service/message"
)

type serviceProvider struct {
	config *config.Config

	dbClient          db.Client
	txManager         db.TxManager
	chatRepository    repository.ChatRepository
	messageRepository repository.MessageRepository

	chatService    service.ChatService
	messageService service.MessageService

	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() *config.Config {
	if s.config == nil {
		cfg, err := config.New()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.Config().Postgres.DSN)
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepo.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = messageRepo.NewRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) MessageService(ctx context.Context) service.MessageService {
	if s.messageService == nil {
		s.messageService = messageService.NewService(
			s.MessageRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.messageService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(
			s.ChatService(ctx),
			s.MessageService(ctx),
		)
	}

	return s.chatImpl
}

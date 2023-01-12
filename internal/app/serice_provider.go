package app

import (
	"context"
	"log"

	"github.com/impopov/note-server-api/internal/config"
	"github.com/impopov/note-server-api/internal/pkg/db"
	"github.com/impopov/note-server-api/internal/repository"
	"github.com/impopov/note-server-api/internal/service/note"
)

type serviceProvider struct {
	db         db.Client
	configPath string
	config     *config.Config

	//repositories
	noteRepository repository.NoteRepository

	//services
	noteService *note.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig(s.configPath)
		if err != nil {
			log.Fatalf("failde to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

// GetDB getter and initializer for db client
func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("can't connect to b err: %s", err.Error())
		}

		s.db = dbc
	}

	return s.db
}

// GetNoteRepository getter ror repository
func (s *serviceProvider) GetNoteRepository(ctx context.Context) repository.NoteRepository {
	if s.noteRepository == nil {
		s.noteRepository = repository.NewRepository(s.GetDB(ctx))
	}

	return s.noteRepository
}

func (s *serviceProvider) GetNoteService(ctx context.Context) *note.Service {
	if s.noteService == nil {
		s.noteService = note.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}

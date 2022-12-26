package note

import "github.com/impopov/note-server-api/internal/repository"

type Service struct {
	noteRepository repository.NoteRepository
}

func NewService(noteRepository repository.NoteRepository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}

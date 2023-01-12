package note

import (
	"context"
	"log"

	"github.com/impopov/note-server-api/internal/model"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (s *Service) GetListNote(ctx context.Context, req *desc.Empty) ([]*model.Note, error) {
	notes, err := s.noteRepository.GetListNote(ctx, req)
	if err != nil {
		log.Printf("error service")
		return nil, err
	}

	return notes, nil
}

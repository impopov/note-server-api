package note

import (
	"context"

	"github.com/impopov/note-server-api/internal/model"
)

func (s *Service) GetNote(ctx context.Context, id int64) (*model.Note, error) {
	res, err := s.noteRepository.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

package note

import (
	"context"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (s *Service) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) error {
	return s.noteRepository.UpdateNote(ctx, req)
}

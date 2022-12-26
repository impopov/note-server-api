package note

import (
	"context"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (s *Service) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error {
	return s.noteRepository.DeleteNote(ctx, req)
}

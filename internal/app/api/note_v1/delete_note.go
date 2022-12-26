package note_v1

import (
	"context"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (i *Implementation) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	err := i.noteService.DeleteNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.Empty{}, nil
}

package note_v1

import (
	"context"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (i *Implementation) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	err := i.noteService.UpdateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.Empty{}, nil
}

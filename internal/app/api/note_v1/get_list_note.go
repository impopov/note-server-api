package note_v1

import (
	"context"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (i *Implementation) GetListNote(ctx context.Context, req *desc.Empty) (*desc.GetListNoteResponse, error) {
	res, err := i.noteService.GetListNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

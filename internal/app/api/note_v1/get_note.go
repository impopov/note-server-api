package note_v1

import (
	"context"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (i *Implementation) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	res, err := i.noteService.GetNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, err
}

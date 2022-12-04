package note_v1

import (
	"context"
	"fmt"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, in *desc.DeleteNoteRequest) (*desc.DeleteNoteResponse, error) {
	fmt.Println("--------------------------")
	fmt.Println("Note deleted by id", in.GetId())

	return &desc.DeleteNoteResponse{Status: "Note deleted"}, nil
}

package note_v1

import (
	"context"
	"fmt"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (n *Implementation) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	fmt.Println("--------------------------")
	fmt.Println("Note deleted by id", req.GetId())

	return &desc.Empty{}, nil
}

package note_v1

import (
	"context"
	"fmt"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (n *Implementation) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	fmt.Println("--------------------------")
	fmt.Println("Note updated")

	fmt.Println("Id", req.GetNote().GetId())
	fmt.Println("Title", req.GetNote().GetTitle())
	fmt.Println("Text", req.GetNote().GetText())
	fmt.Println("Author", req.GetNote().GetAuthor())

	return &desc.Empty{}, nil
}

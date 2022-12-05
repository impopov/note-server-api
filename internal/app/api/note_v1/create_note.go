package note_v1

import (
	"context"
	"fmt"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (n *Implementation) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	fmt.Println("--------------------------")
	fmt.Println("Note created")

	fmt.Println("Title:", req.GetTitle())
	fmt.Println("Text:", req.GetText())
	fmt.Println("Author:", req.GetAuthor())

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}

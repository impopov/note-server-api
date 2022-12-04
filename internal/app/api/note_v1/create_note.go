package note_v1

import (
	"context"
	"fmt"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, in *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	fmt.Println("--------------------------")
	fmt.Println("Create note")
	fmt.Println("Title:", in.GetTitle())
	fmt.Println("Text:", in.GetText())
	fmt.Println("Author:", in.GetAuthor())

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}

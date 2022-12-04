package note_v1

import (
	"context"
	"fmt"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, in *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Println("--------------------------")
	fmt.Println("Note updated")

	fmt.Println("Id", in.Note.GetId())
	fmt.Println("Title", in.Note.GetTitle())
	fmt.Println("Text", in.Note.GetText())
	fmt.Println("Author", in.Note.GetAuthor())

	return &desc.UpdateNoteResponse{Status: "Note updated"}, nil
}

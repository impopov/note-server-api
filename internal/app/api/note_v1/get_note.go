package note_v1

import (
	"context"
	"fmt"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, in *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("--------------------------")
	fmt.Println("Get note")
	fmt.Println("Id", in.GetId())

	return &desc.GetNoteResponse{Note: &desc.Note{
		Id:     2,
		Title:  "The Hitchhiker's Guide to the Galaxy",
		Text:   "Answer is 42",
		Author: "Douglas Adams",
	}}, nil
}

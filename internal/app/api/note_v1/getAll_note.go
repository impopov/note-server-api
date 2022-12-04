package note_v1

import (
	"context"
	"fmt"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (n *Note) GetAllNote(ctx context.Context, in *desc.GetAllNoteRequest) (*desc.GetAllNoteResponse, error) {
	fmt.Println("--------------------------")
	fmt.Println("Get All notes")

	allNotes := &desc.GetAllNoteResponse{Note: []*desc.Note{{
		Id:     1,
		Title:  "Tom Sawyer and Huckleberry Finn",
		Text:   "Right is right, and wrong is wrong, and a body ain't got no business doing wrong when he ain't ignorant and knows better.",
		Author: "Mark Twain",
	}, {
		Id:     2,
		Title:  "The Hitchhiker's Guide to the Galaxy",
		Text:   "Answer is 42",
		Author: "Douglas Adams",
	}}}

	fmt.Println(allNotes)

	return allNotes, nil
}

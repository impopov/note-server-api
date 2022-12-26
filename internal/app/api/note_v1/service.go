package note_v1

import (
	"github.com/impopov/note-server-api/internal/service/note"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteServiceV1Server

	noteService *note.Service
}

func NewImplementation(noteService *note.Service) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}

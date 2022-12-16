package note_v1

import (
	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteServiceV1Server
}

func NewImplementation() *Implementation {
	return &Implementation{}
}

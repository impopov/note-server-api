package note

import (
	"context"
	"log"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

func (s *Service) GetListNote(ctx context.Context, req *desc.Empty) (*desc.GetListNoteResponse, error) {

	notes, err := s.noteRepository.GetListNote(ctx, req)
	if err != nil {
		log.Printf("error srvice")
		return nil, err
	}

	return notes, nil

}

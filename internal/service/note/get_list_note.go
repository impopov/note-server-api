package note

import (
	"context"
	"log"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetListNote(ctx context.Context, req *desc.Empty) (*desc.GetListNoteResponse, error) {
	var notesPb []*desc.Note

	notes, err := s.noteRepository.GetListNote(ctx, req)
	if err != nil {
		log.Printf("error srvice")
		return nil, err
	}

	for _, note := range notes {
		var notePb desc.Note
		notePb.Id = note.Id
		notePb.Title = note.Title
		notePb.Text = note.Text
		notePb.Author = note.Author

		if note.UpdatedAt.Valid {
			notePb.UpdatedAt = timestamppb.New(note.UpdatedAt.Time)
		}

		notePb.CreatedAt = timestamppb.New(note.CreatedAt)

		notesPb = append(notesPb, &notePb)
	}

	return &desc.GetListNoteResponse{Note: notesPb}, nil
}

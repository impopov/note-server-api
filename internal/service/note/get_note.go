package note

import (
	"context"

	repo "github.com/impopov/note-server-api/internal/repository"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	var res *repo.Note
	var updatedAtPb *timestamppb.Timestamp

	res, err := s.noteRepository.GetNote(ctx, req)
	if err != nil {
		return nil, err
	}

	if res.UpdatedAt.Valid {
		updatedAtPb = timestamppb.New(res.UpdatedAt.Time)
	}

	return &desc.GetNoteResponse{Note: &desc.Note{
		Id:        res.Id,
		Title:     res.Title,
		Text:      res.Text,
		Author:    res.Author,
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: updatedAtPb,
	}}, nil

}

package converter

import (
	"github.com/impopov/note-server-api/internal/model"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToNoteInfo(noteInfo *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:  noteInfo.GetTitle(),
		Text:   noteInfo.GetText(),
		Author: noteInfo.GetAuthor(),
	}
}

func ToDeskNoteInfo(noteInfo *model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:  noteInfo.Title,
		Text:   noteInfo.Text,
		Author: noteInfo.Author,
	}
}

func ToDeskNote(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp

	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:        note.Id,
		Info:      ToDeskNoteInfo(&note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToDeskNoteList(notes []*model.Note) []*desc.Note {
	var notesPb []*desc.Note

	for _, note := range notes {
		notePb := new(desc.Note)
		noteInfo := new(desc.NoteInfo)
		notePb.Info = noteInfo

		notePb.Id = note.Id

		notePb.Info.Title = note.Info.Title
		notePb.Info.Text = note.Info.Text
		notePb.Info.Author = note.Info.Author

		notePb.CreatedAt = timestamppb.New(note.CreatedAt)

		if note.UpdatedAt.Valid {
			notePb.UpdatedAt = timestamppb.New(note.UpdatedAt.Time)
		}

		notesPb = append(notesPb, notePb)
	}

	return notesPb
}

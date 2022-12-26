package note_v1

import (
	"context"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
)

const (
	noteTable  = "note"
	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
)

func (i *Implementation) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	res, err := i.noteService.CreateNote(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

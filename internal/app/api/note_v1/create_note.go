package note_v1

import (
	"context"
	"fmt"

	desc "github.com/impopov/note-server-api/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
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

func (n *Implementation) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	fmt.Println("--------------------------")
	fmt.Println("Note created")

	fmt.Println("Title:", req.GetTitle())
	fmt.Println("Text:", req.GetText())
	fmt.Println("Author:", req.GetAuthor())

	dbDSN := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%ssslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDSN)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}

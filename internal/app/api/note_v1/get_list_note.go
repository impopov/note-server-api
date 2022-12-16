package note_v1

import (
	"context"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (n *Implementation) GetListNote(ctx context.Context, req *desc.Empty) (*desc.GetListNoteResponse, error) {
	dbDSN := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDSN)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		From(noteTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*desc.Note
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		var note desc.Note

		err = rows.Scan(
			&note.Id,
			&note.Title,
			&note.Text,
			&note.Author,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		note.CreatedAt = timestamppb.New(createdAt)
		note.CreatedAt = timestamppb.New(updatedAt)

		notes = append(notes, &note)
	}

	return &desc.GetListNoteResponse{Note: notes}, nil
}

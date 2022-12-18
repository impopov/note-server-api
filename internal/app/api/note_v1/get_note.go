package note_v1

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/jmoiron/sqlx"
)

func (n *Implementation) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
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
		PlaceholderFormat(sq.Dollar).
		From(noteTable).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row := db.QueryRowContext(ctx, query, args...)

	var id int64
	var title string
	var text string
	var author string
	var createdAt time.Time
	var updatedAt sql.NullTime
	var updatedAtPb *timestamppb.Timestamp

	err = row.Scan(
		&id,
		&title,
		&text,
		&author,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		updatedAtPb = timestamppb.New(updatedAt.Time)
	}

	return &desc.GetNoteResponse{Note: &desc.Note{
		Id:        id,
		Title:     title,
		Text:      text,
		Author:    author,
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: updatedAtPb,
	}}, nil
}

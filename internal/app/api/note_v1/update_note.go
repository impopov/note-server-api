package note_v1

import (
	"context"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/impopov/note-server-api/pkg/note_v1"

	"github.com/jmoiron/sqlx"
)

func (n *Implementation) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.Empty, error) {
	dbDSN := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDSN)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Update(noteTable).PlaceholderFormat(sq.Dollar).SetMap(sq.Eq{
		"title":      req.GetNote().GetTitle(),
		"text":       req.GetNote().GetText(),
		"author":     req.GetNote().GetAuthor(),
		"updated_at": time.Now(),
	}).Where(sq.Eq{"id": req.GetNote().GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	res, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows != 1 {
		log.Fatalf("expected single row affected, got %d rows affected", rows)
		return nil, err
	}

	return &desc.Empty{}, nil
}

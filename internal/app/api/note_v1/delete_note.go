package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/impopov/note-server-api/pkg/note_v1"

	"github.com/jmoiron/sqlx"
)

func (n *Implementation) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	dbDSN := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDSN)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Delete(noteTable).PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": req.GetId()})
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_ = db.QueryRowContext(ctx, query, args...)

	return &desc.Empty{}, nil
}

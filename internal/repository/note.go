package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/impopov/note-server-api/internal/repository/table"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error)
	GetNote(ctx context.Context, req *desc.GetNoteRequest) (*Note, error)
	GetListNote(ctx context.Context, req *desc.Empty) ([]*Note, error)
	UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) error
	DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) NoteRepository {
	return &repository{
		db: db,
	}
}

type Note struct {
	Id        int64
	Title     string
	Text      string
	Author    string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

func (r *repository) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(req.GetTitle(), req.GetText(), req.GetAuthor()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*Note, error) {
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var note Note
	r.db.QueryRowContext(ctx, query, args...).Scan(
		&note.Id,
		&note.Title,
		&note.Text,
		&note.Author,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	return &note, nil
}

func (r *repository) GetListNote(ctx context.Context, req *desc.Empty) ([]*Note, error) {
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		From(table.Note)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*Note
	for rows.Next() {
		var note Note

		err = rows.Scan(
			&note.Id,
			&note.Title,
			&note.Text,
			&note.Author,
			&note.CreatedAt,
			&note.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		notes = append(notes, &note)
	}

	return notes, nil
}

func (r *repository) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) error {
	builder := sq.Update(table.Note).PlaceholderFormat(sq.Dollar).SetMap(sq.Eq{
		"title":      req.GetNote().GetTitle(),
		"text":       req.GetNote().GetText(),
		"author":     req.GetNote().GetAuthor(),
		"updated_at": time.Now(),
	}).Where(sq.Eq{"id": req.GetNote().GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}

	return nil
}

func (r *repository) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error {
	builder := sq.Delete(table.Note).PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": req.GetId()})
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}

	return nil
}

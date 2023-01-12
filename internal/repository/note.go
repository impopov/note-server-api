package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/impopov/note-server-api/internal/model"
	"github.com/impopov/note-server-api/internal/pkg/db"
	"github.com/impopov/note-server-api/internal/repository/table"
	desc "github.com/impopov/note-server-api/pkg/note_v1"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, noteInfo *model.NoteInfo) (int64, error)
	GetNote(ctx context.Context, id int64) (*model.Note, error)
	GetListNote(ctx context.Context, req *desc.Empty) ([]*model.Note, error)
	UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) error
	DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) error
}

type repository struct {
	client db.Client
}

func NewRepository(client db.Client) NoteRepository {
	return &repository{
		client: client,
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

func (r *repository) CreateNote(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(noteInfo.Title, noteInfo.Text, noteInfo.Author).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "CreateNote",
		QueryRow: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
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

func (r *repository) GetNote(ctx context.Context, id int64) (*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetNote",
		QueryRow: query,
	}

	var note model.Note

	err = r.client.DB().GetContext(ctx, &note, q, args...)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *repository) GetListNote(ctx context.Context, req *desc.Empty) ([]*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		From(table.Note)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetListNote",
		QueryRow: query,
	}

	var notes []*model.Note

	err = r.client.DB().SelectContext(ctx, &notes, q, args...)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *repository) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) error {
	builder := sq.Update(table.Note).
		PlaceholderFormat(sq.Dollar).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": req.GetId()})

	if req.GetNote().GetTitle() != nil {
		builder.Set("title", req.GetNote().GetTitle())
	}

	if req.GetNote().GetTitle() != nil {
		builder.Set("text", req.GetNote().GetText())
	}

	if req.GetNote().GetTitle() != nil {
		builder.Set("author", req.GetNote().GetAuthor())
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "UpdateNote",
		QueryRow: query,
	}

	res, err := r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	rows := res.RowsAffected()
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

	q := db.Query{
		Name:     "DeleteNote",
		QueryRow: query,
	}

	res, err := r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	rows := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}

	return nil
}

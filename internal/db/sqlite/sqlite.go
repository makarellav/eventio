package sqlite

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
)

type DB struct {
	db *sql.DB
}

//go:embed schema.sql
var schema string

func New(ctx context.Context) (*DB, error) {
	db, err := sql.Open("sqlite", "./eventio.db")

	if err != nil {
		return nil, fmt.Errorf("db.new: %w", err)
	}

	_, err = db.ExecContext(ctx, schema)

	if err != nil {
		return nil, fmt.Errorf("db.new: %w", err)
	}

	return &DB{db: db}, nil
}

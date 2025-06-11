package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBTX interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryxContext(context.Context, string, ...any) (*sqlx.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
	QueryRowxContext(context.Context, string, ...any) *sqlx.Row
}

type DataStore interface {
	Atomic(ctx context.Context, fn func(DataStore) error) error
	UserRepository() UserRepository
}

type dataStore struct {
	conn *sqlx.DB
	db   DBTX
}

func NewDataStore(db *sqlx.DB) DataStore {
	return &dataStore{
		conn: db,
		db:   db,
	}
}

func (ds *dataStore) Atomic(ctx context.Context, fn func(DataStore) error) error {
	tx, err := ds.conn.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	if err := fn(&dataStore{conn: ds.conn, db: tx}); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (ds *dataStore) UserRepository() UserRepository {
	return NewUserRepository(ds.db)
}

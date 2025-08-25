package db

import "github.com/jackc/pgx/v5/pgxpool"

type Store interface {
	Querier
	// Add other methods or transactions here if needed
}

type SqlStore struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) Store {
	return &SqlStore{
		Queries: New(db),
		db:      db,
	}
}

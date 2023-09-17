package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Provide all functions for db queries (Queries) and transactions (db)
type Library struct {
	*Queries
	db *pgxpool.Pool
}

func NewLibrary(connPool *pgxpool.Pool) Library {
	return Library{
		db:      connPool,
		Queries: New(connPool),
	}
}

package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Library interface {
	Querier
}

// Provide all functions for db queries (Queries) and transactions (db)
type SQLLibrary struct {
	*Queries
	db *pgxpool.Pool
}

func NewLibrary(connPool *pgxpool.Pool) Library {
	return &SQLLibrary{
		db:      connPool,
		Queries: New(connPool),
	}
}

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Library interface {
	Querier
	BorrowTx(ctx context.Context, arg BorrowTxParams) (BorrowTxResult, error)
	CreateMemberTx(ctx context.Context, arg CreateMemberTxParams) (CreateMemberTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) error
	ProcessBorrowTx(ctx context.Context, arg ProcessBorrowTxParams) (Transaction, error)
	ProcessReturnTx(ctx context.Context, arg ProcessReturnTxParams) (Transaction, error)
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

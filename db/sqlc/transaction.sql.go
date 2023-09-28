// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: transaction.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
  member_id,
  borrow_id,
  purpose,
  status
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, member_id, admin_id, borrow_id, purpose, status, note, created_at, updated_at, deleted_at
`

type CreateTransactionParams struct {
	MemberID uuid.UUID `json:"member_id"`
	BorrowID uuid.UUID `json:"borrow_id"`
	Purpose  Purpose   `json:"purpose"`
	Status   Status    `json:"status"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.MemberID,
		arg.BorrowID,
		arg.Purpose,
		arg.Status,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.MemberID,
		&i.AdminID,
		&i.BorrowID,
		&i.Purpose,
		&i.Status,
		&i.Note,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, member_id, admin_id, borrow_id, purpose, status, note, created_at, updated_at, deleted_at FROM transactions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransaction(ctx context.Context, id uuid.UUID) (Transaction, error) {
	row := q.db.QueryRow(ctx, getTransaction, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.MemberID,
		&i.AdminID,
		&i.BorrowID,
		&i.Purpose,
		&i.Status,
		&i.Note,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getTransactionAndBorrowDetail = `-- name: GetTransactionAndBorrowDetail :one
SELECT t.id trx_id, t.member_id trx_member_id, bd.id bd_id FROM transactions t
JOIN borrow_details bd ON t.borrow_id = bd.id
WHERE t.id = $1 LIMIT 1
`

type GetTransactionAndBorrowDetailRow struct {
	TrxID       uuid.UUID `json:"trx_id"`
	TrxMemberID uuid.UUID `json:"trx_member_id"`
	BdID        uuid.UUID `json:"bd_id"`
}

func (q *Queries) GetTransactionAndBorrowDetail(ctx context.Context, id uuid.UUID) (GetTransactionAndBorrowDetailRow, error) {
	row := q.db.QueryRow(ctx, getTransactionAndBorrowDetail, id)
	var i GetTransactionAndBorrowDetailRow
	err := row.Scan(&i.TrxID, &i.TrxMemberID, &i.BdID)
	return i, err
}

const updateTransaction = `-- name: UpdateTransaction :one
UPDATE transactions 
SET 
  admin_id=$1,
  status=$2,
  note=$3
WHERE id=$4
RETURNING id, member_id, admin_id, borrow_id, purpose, status, note, created_at, updated_at, deleted_at
`

type UpdateTransactionParams struct {
	AdminID pgtype.UUID `json:"admin_id"`
	Status  Status      `json:"status"`
	Note    pgtype.Text `json:"note"`
	ID      uuid.UUID   `json:"id"`
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, updateTransaction,
		arg.AdminID,
		arg.Status,
		arg.Note,
		arg.ID,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.MemberID,
		&i.AdminID,
		&i.BorrowID,
		&i.Purpose,
		&i.Status,
		&i.Note,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type BorrowTxParams struct {
	CreateBorrowParams
	CreateTransactionParams
	Quantity int
}

type BorrowTxResult struct {
	ID        uuid.UUID   `json:"id"`
	MemberID  uuid.UUID   `json:"member_id"`
	AdminID   pgtype.UUID `json:"admin_id"`
	BorrowID  uuid.UUID   `json:"borrow_id"`
	Purpose   Purpose     `json:"purpose"`
	Status    Status      `json:"status"`
	Note      pgtype.Text `json:"note"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (library *SQLLibrary) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := library.db.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}

func (library *SQLLibrary) BorrowTx(ctx context.Context, arg BorrowTxParams) (BorrowTxResult, error) {
	var result BorrowTxResult
	err := library.execTx(ctx, func(q *Queries) error {
		// Create entry in borrow_details table
		borrowDetail, err := q.CreateBorrow(ctx, arg.CreateBorrowParams)
		if err != nil {
			return err
		}
		arg.CreateTransactionParams.BorrowID = borrowDetail.ID
		// Create entry in transactions table
		transaction, err := q.CreateTransaction(ctx, arg.CreateTransactionParams)
		result = BorrowTxResult{
			ID:        transaction.ID,
			MemberID:  transaction.MemberID,
			AdminID:   transaction.AdminID,
			BorrowID:  transaction.BorrowID,
			Purpose:   transaction.Purpose,
			Status:    transaction.Status,
			Note:      transaction.Note,
			CreatedAt: transaction.CreatedAt,
			UpdatedAt: transaction.UpdatedAt,
		}
		if err != nil {
			return err
		}
		// Update entry in books table
		_, err = q.UpdateBook(ctx, UpdateBookParams{
			ID: borrowDetail.BookID,
			Quantity: pgtype.Int4{
				Int32: int32(arg.Quantity) - 1,
				Valid: true,
			},
		})
		return err
	})

	return result, err
}

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type BorrowTxParams struct {
	CreateBorrowParams
	CreateTransactionParams
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
		// Find books in books table, check quantity. Locking the row with SELECT FOR UPDATE
		book, err := q.GetBookForUpdate(ctx, arg.CreateBorrowParams.BookID)
		if err != nil {
			return fmt.Errorf("failed to get book for update: %w", err)
		}

		if book.Quantity == 0 {
			return errors.New("books are not available to borrow")
		}

		// Create entry in borrow_details table
		borrowDetail, err := q.CreateBorrow(ctx, arg.CreateBorrowParams)
		if err != nil {
			return fmt.Errorf("failed to create borrow detail entry: %w", err)
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
			return fmt.Errorf("failed to create transaction entry: %w", err)

		}
		// Update entry in books table
		_, err = q.UpdateBook(ctx, UpdateBookParams{
			ID: borrowDetail.BookID,
			Quantity: pgtype.Int4{
				Int32: int32(book.Quantity) - 1,
				Valid: true,
			},
		})
		return err
	})

	return result, err
}

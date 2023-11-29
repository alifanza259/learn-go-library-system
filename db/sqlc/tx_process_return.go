package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type ProcessReturnTxParams struct {
	UpdateTransactionParams
	Transaction GetTransactionAssociatedDetailRow
	AfterUpdate func(transaction GetTransactionAssociatedDetailRow, status Status, note string) error
}

func (library *SQLLibrary) ProcessReturnTx(ctx context.Context, arg ProcessReturnTxParams) (Transaction, error) {
	var result Transaction
	err := library.execTx(ctx, func(q *Queries) error {
		var err error

		result, err = q.UpdateTransaction(ctx, arg.UpdateTransactionParams)
		if err != nil {
			return err
		}
		if arg.UpdateTransactionParams.Status == StatusApproved {
			book, err := q.GetBook(ctx, arg.Transaction.BID)
			if err != nil {
				return err
			}

			_, err = q.UpdateBook(ctx, UpdateBookParams{
				ID: book.ID,
				Quantity: pgtype.Int4{
					Int32: book.Quantity + 1,
					Valid: true,
				},
			})
			if err != nil {
				return err
			}
		}

		if err := arg.AfterUpdate(arg.Transaction, arg.UpdateTransactionParams.Status, arg.UpdateTransactionParams.Note.String); err != nil {
			return err
		}

		return nil
	})

	return result, err
}

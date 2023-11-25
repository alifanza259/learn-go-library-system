package db

import (
	"context"
)

type ProcessBorrowTxParams struct {
	UpdateTransactionParams
	Transaction GetTransactionAssociatedDetailRow
	AfterUpdate func(transaction GetTransactionAssociatedDetailRow, status Status, note string) error
}

func (library *SQLLibrary) ProcessBorrowTx(ctx context.Context, arg ProcessBorrowTxParams) (Transaction, error) {
	var result Transaction
	err := library.execTx(ctx, func(q *Queries) error {
		var err error

		result, err = q.UpdateTransaction(ctx, arg.UpdateTransactionParams)
		if err != nil {
			return err
		}

		if err := arg.AfterUpdate(arg.Transaction, arg.UpdateTransactionParams.Status, arg.UpdateTransactionParams.Note.String); err != nil {
			return err
		}

		return nil
	})

	return result, err
}

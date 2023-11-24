package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type VerifyEmailTxParams struct {
	EmailVerif EmailVerification
}

func (library *SQLLibrary) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) error {
	err := library.execTx(ctx, func(q *Queries) error {
		err := library.UpdateEmailVerification(ctx, UpdateEmailVerificationParams{
			Token:  arg.EmailVerif.Token,
			IsUsed: true,
		})
		if err != nil {
			return err
		}

		_, err = library.UpdateMember(ctx, UpdateMemberParams{
			EmailVerifiedAt: pgtype.Timestamptz{
				Time:  time.Now(),
				Valid: true,
			},
			ID: arg.EmailVerif.MemberID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

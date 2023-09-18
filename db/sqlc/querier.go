// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error)
	GetAdmin(ctx context.Context, id uuid.UUID) (Admin, error)
	GetMember(ctx context.Context, id uuid.UUID) (Member, error)
	GetMemberByEmail(ctx context.Context, email string) (Member, error)
	ListAdmin(ctx context.Context) ([]ListAdminRow, error)
	ListMembers(ctx context.Context) ([]ListMembersRow, error)
}

var _ Querier = (*Queries)(nil)

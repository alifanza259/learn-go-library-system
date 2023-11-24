// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: member.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createMember = `-- name: CreateMember :one
INSERT INTO members (
  email, 
  first_name, 
  last_name, 
  dob,
  gender,
  password
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, email, first_name, last_name, dob, gender, password, password_changed_at, last_accessed_at, created_at, updated_at, deleted_at, email_verified_at
`

type CreateMemberParams struct {
	Email     string      `json:"email"`
	FirstName string      `json:"first_name"`
	LastName  pgtype.Text `json:"last_name"`
	Dob       pgtype.Date `json:"dob"`
	Gender    Gender      `json:"gender"`
	Password  string      `json:"password"`
}

func (q *Queries) CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error) {
	row := q.db.QueryRow(ctx, createMember,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.Dob,
		arg.Gender,
		arg.Password,
	)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Dob,
		&i.Gender,
		&i.Password,
		&i.PasswordChangedAt,
		&i.LastAccessedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.EmailVerifiedAt,
	)
	return i, err
}

const getMember = `-- name: GetMember :one
SELECT id, email, first_name, last_name, dob, gender, password, password_changed_at, last_accessed_at, created_at, updated_at, deleted_at, email_verified_at FROM members
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMember(ctx context.Context, id uuid.UUID) (Member, error) {
	row := q.db.QueryRow(ctx, getMember, id)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Dob,
		&i.Gender,
		&i.Password,
		&i.PasswordChangedAt,
		&i.LastAccessedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.EmailVerifiedAt,
	)
	return i, err
}

const getMemberByEmail = `-- name: GetMemberByEmail :one
SELECT id, email, first_name, last_name, dob, gender, password, password_changed_at, last_accessed_at, created_at, updated_at, deleted_at, email_verified_at FROM members
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetMemberByEmail(ctx context.Context, email string) (Member, error) {
	row := q.db.QueryRow(ctx, getMemberByEmail, email)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Dob,
		&i.Gender,
		&i.Password,
		&i.PasswordChangedAt,
		&i.LastAccessedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.EmailVerifiedAt,
	)
	return i, err
}

const listMembers = `-- name: ListMembers :many
SELECT 
  id,
  email,
  first_name,
  last_name,
  dob,
  gender,
  created_at,
  updated_at,
  deleted_at
FROM members
ORDER BY first_name
`

type ListMembersRow struct {
	ID        uuid.UUID          `json:"id"`
	Email     string             `json:"email"`
	FirstName string             `json:"first_name"`
	LastName  pgtype.Text        `json:"last_name"`
	Dob       pgtype.Date        `json:"dob"`
	Gender    Gender             `json:"gender"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
}

func (q *Queries) ListMembers(ctx context.Context) ([]ListMembersRow, error) {
	rows, err := q.db.Query(ctx, listMembers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListMembersRow{}
	for rows.Next() {
		var i ListMembersRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.Dob,
			&i.Gender,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

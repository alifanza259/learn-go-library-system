// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Gender string

const (
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
	GenderUnknown Gender = "unknown"
)

func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

type NullGender struct {
	Gender Gender `json:"gender"`
	Valid  bool   `json:"valid"` // Valid is true if Gender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGender) Scan(value interface{}) error {
	if value == nil {
		ns.Gender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Gender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Gender), nil
}

func (e Gender) Valid() bool {
	switch e {
	case GenderMale,
		GenderFemale,
		GenderUnknown:
		return true
	}
	return false
}

type Admin struct {
	ID             uuid.UUID          `json:"id"`
	Email          string             `json:"email"`
	FirstName      string             `json:"first_name"`
	LastName       pgtype.Text        `json:"last_name"`
	Password       string             `json:"password"`
	Permission     string             `json:"permission"`
	LastAccessedAt pgtype.Timestamptz `json:"last_accessed_at"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	DeletedAt      pgtype.Timestamptz `json:"deleted_at"`
}

type Member struct {
	ID                uuid.UUID          `json:"id"`
	Email             string             `json:"email"`
	FirstName         string             `json:"first_name"`
	LastName          pgtype.Text        `json:"last_name"`
	Dob               pgtype.Date        `json:"dob"`
	Gender            Gender             `json:"gender"`
	Password          string             `json:"password"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	LastAccessedAt    pgtype.Timestamptz `json:"last_accessed_at"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	DeletedAt         pgtype.Timestamptz `json:"deleted_at"`
}
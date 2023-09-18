package token

import db "github.com/alifanza259/learn-go-library-system/db/sqlc"

type Maker interface {
	CreateToken(member db.Member) (string, int, error)
	VerifyToken(tokenString string) (*Payload, error)
}

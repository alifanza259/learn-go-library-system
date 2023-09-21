package token

import (
	"time"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
)

type Maker interface {
	// duration is passed here as parameter for unit test purposes (test unhappy case)
	CreateToken(member db.Member, duration time.Duration) (string, int, error)
	VerifyToken(tokenString string) (*Payload, error)
}

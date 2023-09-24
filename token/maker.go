package token

import (
	"time"
)

type Maker interface {
	// duration is passed here as parameter for unit test purposes (test unhappy case)
	CreateToken(email string, duration time.Duration, purpose string) (string, int, error)
	VerifyToken(tokenString string, purpose string) (*Payload, error)
}

package token

import (
	"fmt"
	"time"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateToken(member db.Member) (string, int, error) {
	expiresAt := time.Now().Add(time.Second * 15).UnixMilli()
	claims := MyCustomClaims{
		member.ID.String(),
		member.Email,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "login",
		},
	}

	mySigningKey := []byte("AllYourBase")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", 0, fmt.Errorf("failed to create token: %w", err)
	}

	return tokenString, int(expiresAt), nil
}

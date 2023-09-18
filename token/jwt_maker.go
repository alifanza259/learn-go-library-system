package token

import (
	"errors"
	"fmt"
	"time"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/dgrijalva/jwt-go"
)

type JWTMaker struct {
	secret   string
	duration int
}

type Payload struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewJWTMaker(secret string, duration int) Maker {
	return &JWTMaker{
		secret:   secret,
		duration: duration,
	}
}

func (maker *JWTMaker) CreateToken(member db.Member) (string, int, error) {
	expiresAt := time.Now().Add(time.Second * time.Duration(maker.duration)).Unix()
	claims := Payload{
		member.ID.String(),
		member.Email,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiresAt,
			Issuer:    "login",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString([]byte(maker.secret))
	if err != nil {
		return "", 0, fmt.Errorf("failed to create token: %w", err)
	}

	return tokenString, int(expiresAt), nil
}

func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token is invalid")
		}

		return []byte(maker.secret), nil
	})

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errors.New("token has expired")) {
			return nil, errors.New("token has expired")
		}
		return nil, errors.New("token is invalid")
	}

	payload, ok := token.Claims.(*Payload)
	if !ok {
		return nil, errors.New("token is invalid")
	}
	return payload, nil
}

package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTMaker struct {
	secret      string
	secretAdmin string
}

type Payload struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewJWTMaker(secret string, secretAdmin string) Maker {
	return &JWTMaker{
		secret:      secret,
		secretAdmin: secretAdmin,
	}
}

func (maker *JWTMaker) CreateToken(email string, duration time.Duration, purpose string) (string, int, error) {
	expiresAt := time.Now().Add(duration).Unix()
	claims := Payload{
		email,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiresAt,
			Issuer:    "login",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	var secret string
	if purpose == "admin" {
		secret = maker.secretAdmin
	} else {
		secret = maker.secret
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, fmt.Errorf("failed to create token: %w", err)
	}

	return tokenString, int(expiresAt), nil
}

func (maker *JWTMaker) VerifyToken(tokenString string, purpose string) (*Payload, error) {
	var secret string
	if purpose == "admin" {
		secret = maker.secretAdmin
	} else {
		secret = maker.secret
	}

	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token is invalid")
		}

		return []byte(secret), nil
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

package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secret      string
	secretAdmin string
}

type Payload struct {
	Email     string    `json:"email"`
	ID        string    `json:"id"`
	Issuer    string    `json:"iss"`
	ExpiresAt time.Time `json:"exp"`
	IssuedAt  time.Time `json:"iat"`
}

type CustomClaimsGolangJwt struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.RegisteredClaims
}

func NewJWTMaker(secret string, secretAdmin string) Maker {
	return &JWTMaker{
		secret:      secret,
		secretAdmin: secretAdmin,
	}
}

// TODO: change purpose to enum/const
func (maker *JWTMaker) CreateToken(email string, id string, duration time.Duration, purpose string) (string, int, error) {
	expiresAt := time.Now().Add(duration)
	claims := &CustomClaimsGolangJwt{
		email,
		id,
		jwt.RegisteredClaims{
			Issuer:    "login",
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var secret string
	if purpose == "admin" {
		secret = maker.secretAdmin
	} else {
		secret = maker.secret
	}

	tokenString, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", 0, fmt.Errorf("failed to create token: %w", err)
	}

	return tokenString, int(expiresAt.Unix()), nil

}

func (maker *JWTMaker) VerifyToken(tokenString string, purpose string) (*Payload, error) {
	var secret string
	if purpose == "admin" {
		secret = maker.secretAdmin
	} else {
		secret = maker.secret
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsGolangJwt{}, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token is invalid")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, errors.New("token is invalid")
	}

	claims, ok := token.Claims.(*CustomClaimsGolangJwt)
	if !ok {
		return nil, errors.New("token is invalid")
	}

	payload := &Payload{
		Email:     claims.Email,
		ID:        claims.ID,
		Issuer:    claims.Issuer,
		ExpiresAt: claims.ExpiresAt.Time,
		IssuedAt:  claims.IssuedAt.Time,
	}
	return payload, nil
}

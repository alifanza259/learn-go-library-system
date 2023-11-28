package token

import (
	"encoding/json"
	"time"

	"aidanwoods.dev/go-paseto"
)

type PasetoMaker struct {
	symmetricKey      paseto.V4SymmetricKey
	symmetricKeyAdmin paseto.V4SymmetricKey
}

func NewPasetoMaker(symmetricKey string, symmetricKeyAdmin string) (Maker, error) {
	key, err := paseto.V4SymmetricKeyFromBytes([]byte(symmetricKey))
	if err != nil {
		return nil, err
	}
	keyAdmin, err := paseto.V4SymmetricKeyFromBytes([]byte(symmetricKeyAdmin))
	if err != nil {
		return nil, err
	}
	return &PasetoMaker{
		symmetricKey:      key,
		symmetricKeyAdmin: keyAdmin,
	}, nil
}

func (maker *PasetoMaker) CreateToken(email string, id string, duration time.Duration, purpose string) (string, int, error) {
	token := paseto.NewToken()
	expiredAt := time.Now().Add(duration)
	token.SetExpiration(expiredAt)
	token.SetString("email", email)
	token.SetString("id", id)
	token.SetIssuer("login")
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())

	symKey := maker.symmetricKey
	if purpose == "admin" {
		symKey = maker.symmetricKeyAdmin
	}
	encrypted := token.V4Encrypt(symKey, nil)

	return encrypted, int(expiredAt.Unix()), nil
}

func (maker *PasetoMaker) VerifyToken(tokenString string, purpose string) (*Payload, error) {
	symKey := maker.symmetricKey
	if purpose == "admin" {
		symKey = maker.symmetricKeyAdmin
	}

	parser := paseto.NewParserForValidNow()
	token, err := parser.ParseV4Local(symKey, tokenString, nil)
	if err != nil {
		return nil, err
	}

	var payload *Payload
	err = json.Unmarshal(token.ClaimsJSON(), &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

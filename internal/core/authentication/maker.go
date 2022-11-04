package authentication

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	domainErrors "github.com/kenethrrizzo/skeleton-auth/internal/core/errors"
)

// interface that defines methods to use
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be a least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
    payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, domainErrors.ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)	
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, domainErrors.ErrExpiredToken) {
			return nil, domainErrors.ErrExpiredToken
		}

		return nil, domainErrors.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, domainErrors.ErrInvalidToken
	}

	return payload, nil
}

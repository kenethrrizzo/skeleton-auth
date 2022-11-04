package authentication_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	auth "github.com/kenethrrizzo/skeleton-auth/internal/core/authentication"
	domainErrors "github.com/kenethrrizzo/skeleton-auth/internal/core/errors"
	"github.com/stretchr/testify/require"
)

const randomSecretKey = "98ahbnah823la4532so182hfbfbgsd1214"

func TestJWTMaker(t *testing.T) {
	maker, err := auth.NewJWTMaker(randomSecretKey)
	require.NoError(t, err)

	username := "Keneth"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)

	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWT(t *testing.T) {
	maker, err := auth.NewJWTMaker(randomSecretKey)
	require.NoError(t, err)

	token, err := maker.CreateToken("Keneth", -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, domainErrors.ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTAlgNone(t *testing.T) {
	payload, err := auth.NewPayload("Keneth", time.Minute)
	require.NoError(t, err)

	token := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	tokenStr, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := auth.NewJWTMaker(randomSecretKey)
	require.NoError(t, err)

	payload, err = maker.VerifyToken(tokenStr)
	require.Error(t, err)
	require.EqualError(t, err, domainErrors.ErrInvalidToken.Error())
	require.Nil(t, payload)
}

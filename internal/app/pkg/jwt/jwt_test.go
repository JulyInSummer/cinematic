package jwt_test

import (
	"github.com/JulyInSummer/cinematic/internal/app/pkg/jwt"
	"github.com/stretchr/testify/require"
	"testing"
)

func createToken(email, secret string) (string, error) {
	return jwt.CreateToken(email, secret)
}

func TestCreateToken(t *testing.T) {
	email := "test@test.com"
	secret := "secret"

	tokenStr, err := createToken(email, secret)
	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)
}

func TestVerifyTokenSuccess(t *testing.T) {
	secret := "secret"
	token, err := createToken("test@test.com", secret)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	err = jwt.VerifyToken(token, secret)
	require.NoError(t, err)
}

func TestVerifyTokenFail(t *testing.T) {
	secret := "secret"
	token := "dsdhfdsjfsmldfcdbvehxnfmerbfbvxdgvfxfvx"

	err := jwt.VerifyToken(token, secret)
	require.Error(t, err)
}

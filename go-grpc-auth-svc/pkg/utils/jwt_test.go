package utils

import (
	"testing"

	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	secretKey := "secret"
	issuer := "go-grpc-auth-svc"
	expirationHours := int64(24)

	user := models.User{
		Id:    1,
		Email: "test@example.com",
	}

	jwtWrapper := JwtWrapper{
		SecretKey:       secretKey,
		Issuer:          issuer,
		ExpirationHours: expirationHours,
	}

	signedToken, err := jwtWrapper.GenerateToken(user)
	require.NoError(t, err)
	require.NotEmpty(t, signedToken)
}

func TestValidateToken(t *testing.T) {
	secretKey := "secret"
	issuer := "go-grpc-auth-svc"
	expirationHours := int64(24)

	jwtWrapper := JwtWrapper{
		SecretKey:       secretKey,
		Issuer:          issuer,
		ExpirationHours: expirationHours,
	}

	user := models.User{
		Id:    1,
		Email: "test@example.com",
	}
	signedToken, err := jwtWrapper.GenerateToken(user)
	require.NoError(t, err)

	claims, err := jwtWrapper.ValidateToken(signedToken)
	require.NoError(t, err)
	require.NotNil(t, claims)
	require.Equal(t, user.Id, claims.Id)
	require.Equal(t, user.Email, claims.Email)
}

func TestExpiredTokenValidation(t *testing.T) {
	secretKey := "secret"
	issuer := "go-grpc-auth-svc"
	expirationHours := int64(-1)

	jwtWrapper := JwtWrapper{
		SecretKey:       secretKey,
		Issuer:          issuer,
		ExpirationHours: expirationHours,
	}

	user := models.User{
		Id:    1,
		Email: "test@example.com",
	}
	signedToken, err := jwtWrapper.GenerateToken(user)
	require.NoError(t, err)

	claims, err := jwtWrapper.ValidateToken(signedToken)
	require.Error(t, err)
	require.Nil(t, claims)
	require.Contains(t, err.Error(), "token is expired by 1h0m0s")
}

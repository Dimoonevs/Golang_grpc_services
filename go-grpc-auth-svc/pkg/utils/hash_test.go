package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHash(t *testing.T) {

	password := "qwerty"
	hash := HashPassword(password)

	require.NotNil(t, hash)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	require.NoError(t, err)
}

func TestCheck(t *testing.T) {

	password := "qwerty"
	hash := HashPassword(password)
	isPaswordCorrect := CheckPasswordHash(password, hash)

	require.True(t, isPaswordCorrect)

	incorrectPassword := "qwerty123"
	isPaswordIncorrect := CheckPasswordHash(incorrectPassword, hash)

	require.False(t, isPaswordIncorrect)
}

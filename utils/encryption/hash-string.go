package encryption

import (
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	ErrCouldNotSaltPassword = "ErrCouldNotSaltPassword: Failed to encrypt the given password"
)

func HashString(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Error(ErrCouldNotSaltPassword)
	}

	salt := string(hash)

	return salt, nil
}
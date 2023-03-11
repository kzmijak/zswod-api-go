package encryption

import (
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	ErrHashNotEqual = "ErrHasNotEqual: Tested string failed hash equality check"
)

func CompareWithHash(testedString string, hashedString string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(testedString)); err != nil {
		return errors.Error(ErrHashNotEqual)
	}

	return nil
}
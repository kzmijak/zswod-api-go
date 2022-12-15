package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrUnexpectedSigningMethod = "err_unexpected_signing_method: Unexpected signing method"
)

func (s JwtService) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Error(ErrUnexpectedSigningMethod)
		}

		return []byte(s.cfg.Secret), nil
	})

	return token, err
}
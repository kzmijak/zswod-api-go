package jwt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

func (s JwtService) GenerateToken(guid string, role role.Role) (string, error) {
	tokenLifespan, err := strconv.Atoi(s.cfg.LifespanHours)

	if err != nil {
		return "", errors.Error(ErrFailedToConvertLifespan).Wrap(err)
	}

	claims := jwt.MapClaims{
		CLAIM_AUTHORIZED: true,
		CLAIM_GUID:       guid,
		CLAIM_ROLE:       role,
		CLAIM_EXPIRED:    time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.cfg.Secret))

	if err != nil {
		return "", errors.Error(ErrFailedSigningToken).Wrap(err)
	}

	return tokenString, nil
}
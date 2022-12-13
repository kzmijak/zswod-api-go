package authentication

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	CLAIM_AUTHORIZED = "authorized"
	CLAIM_EMAIL = "email"
	CLAIM_ROLE = "role"
	CLAIM_EXPIRED = "exp"
)

const (
	ErrFailedParsingToken = "err_failed_parsing_token: Could not parse token to string"
)

func GenerateJwt(email string, role role.Role) (string, error) {
	var signingKey = []byte("dfssdfsd") // TODO
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims[CLAIM_AUTHORIZED] = true
	claims[CLAIM_EMAIL] = email
	claims[CLAIM_ROLE] = role // TODO, make it an array of enums
	claims[CLAIM_EXPIRED] = time.Now().Add(time.Minute * 30).Unix()

	tokenStringified, err := token.SignedString(signingKey)
	if err != nil {
		return "", errors.Error(ErrFailedParsingToken)
	}

	return tokenStringified, nil;
}
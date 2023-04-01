package userService

import (
	"time"

	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrTokenNotFound = "ErrTokenNotFound: Token not found for this user"
	ErrTokenExpired = "ErrTokenExpired: Token is only valid 2 hours from creation"
	ErrTokenHasNoOwner = "ErrTokenHasNoOwner: Token owner not found"
)

func (s UserService) GetResetPasswordTokenOwner(tokenId string, tx *ent.Tx) (*ent.User, error) {
	token, err := s.GetResetPasswordToken(tokenId, tx)
	if err != nil {
		return nil,errors.Error(ErrTokenNotFound)
	}

	twoHoursAfterCreation := token.CreatedAt.Add(time.Hour * 2)

	if (time.Now().After(twoHoursAfterCreation)) {
		return nil,errors.Error(ErrTokenExpired)
	}

	tokenOwner, err := token.QueryOwner().Only(s.ctx)
	if err != nil {
		return nil, errors.Error(ErrTokenHasNoOwner)
	}

	return tokenOwner, nil
}
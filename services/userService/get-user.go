package userService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotParseUserId = "ErrCouldNotParseUserId: Could not parse specified string to UUID"
	ErrUserIdNotFound = "ErrUserIdNotFound: Could not find user by given ID"
)

func (s UserService) GetUser(userId string, tx *ent.Tx) (*ent.User, error) {
	userIdParsed, err := uuid.Parse(userId)
	if err != nil {
		return nil, errors.Error(ErrCouldNotParseUserId)
	}

	user, err := tx.User.Get(s.ctx, userIdParsed)
	if err != nil {
		return nil, errors.Error(ErrUserIdNotFound)
	}

	return user, nil
}
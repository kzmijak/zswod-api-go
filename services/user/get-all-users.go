package user

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotQuery = "ErrCouldNotQuery: Failed to query for users"
)

func (s UserService) GetAllUsers(tx *ent.Tx) ([]*ent.User, error)  {
	if s.ctx == nil {
		return nil, errors.Error(ErrNoContext)
	}

	users, err := tx.User.Query().All(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrCouldNotQuery)
	}

	return users, nil
}
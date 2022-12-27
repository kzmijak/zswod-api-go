package user

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotQuery = "err_could_not_query: Failed to query for users"
)

func (s UserService) GetAllUsers() ([]*ent.User, error)  {
	if s.ctx == nil {
		return nil, errors.Error(ErrNoContext)
	}

	users, err := database.Client.User.Query().All(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrCouldNotQuery)
	}

	return users, nil
}
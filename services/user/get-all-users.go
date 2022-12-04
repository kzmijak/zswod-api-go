package user

import (
	"context"

	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotQuery = "err_could_not_query: Failed to query for users"
)

func GetAllUsers(ctx context.Context) ([]*ent.User, error)  {
	users, err := database.Client.User.Query().All(ctx)

	if err != nil {
		return nil, errors.Error(ErrCouldNotQuery)
	}

	return users, nil
}
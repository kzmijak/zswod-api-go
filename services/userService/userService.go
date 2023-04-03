package userService

import (
	"context"

	"github.com/kzmijak/zswod_api_go/selectors/userSelectors"
)

const (
	ErrNoContext = "ErrNoContext: Context is required to run this method"
	ErrNoJwtService = "ErrNoJwtService: Jwt Service instance is required to run this method"
)

type UserService struct {
	ctx context.Context
	userSelectors userSelectors.UserSelectors
}

func New(ctx context.Context) UserService {
	return UserService{
		ctx: ctx,
		userSelectors: userSelectors.Initialize(ctx),
	}
}


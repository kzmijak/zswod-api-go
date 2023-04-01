package userService

import (
	"context"
)

const (
	ErrNoContext = "ErrNoContext: Context is required to run this method"
	ErrNoJwtService = "ErrNoJwtService: Jwt Service instance is required to run this method"
)

type UserService struct {
	ctx context.Context
}

func New() UserService {
	return UserService{}
}

func (s UserService) WithContext(ctx context.Context) (UserService) {
	s.ctx = ctx
	return s
}

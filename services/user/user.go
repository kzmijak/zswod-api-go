package user

import (
	"context"

	"github.com/kzmijak/zswod_api_go/services/jwt"
)

const (
	ErrNoContext = "ErrNoContext: Context is required to run this method"
	ErrNoJwtService = "ErrNoJwtService: Jwt Service instance is required to run this method"
)

type UserService struct {
	jwtService jwt.JwtService
	ctx context.Context
}

func New() UserService {
	return UserService{}
}

func (s UserService) WithContext(ctx context.Context) (UserService) {
	s.ctx = ctx
	return s
}

func (s UserService) WithJwtService(jwtService jwt.JwtService) (UserService) {
	s.jwtService = jwtService
	return s
}
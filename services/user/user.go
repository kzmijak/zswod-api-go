package user

import (
	"context"

	"github.com/kzmijak/zswod_api_go/services/jwt"
)

const (
	ErrNoContext = "err_no_context: Context is required to run this method"
	ErrNoJwtService = "err_no_jwt_service: Jwt Service instance is required to run this method"
)

type UserService struct {
	jwtService jwt.JwtService
	ctx context.Context
}

func New() UserService {
	return UserService{}
}

func (UserService) WithContext(ctx context.Context) (s UserService) {
	s.ctx = ctx
	return
}

func (UserService) WithJwtService(jwtService jwt.JwtService) (s UserService) {
	s.jwtService = jwtService
	return
}
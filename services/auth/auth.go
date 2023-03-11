package auth

import "context"

type AuthService struct {
	ctx context.Context
}

func New() AuthService {
	return AuthService{}
}

func (s AuthService) WithContext(ctx context.Context) AuthService {
	s.ctx = ctx
	return s
}
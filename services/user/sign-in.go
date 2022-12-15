package user

import (
	"github.com/kzmijak/zswod_api_go/ent/user"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"golang.org/x/crypto/bcrypt"
)

type SignInRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

const (
	ErrUserNotFound = "err_user_not_found: Could not match any user with provided username"
	ErrInvalidPassword = "err_invalid_password: Password for the user is incorrect"
	ErrInvalidClaims = "err_invalid_claims: Could not generate token for this user"
)

func (s UserService) SignIn(request SignInRequest) (string, error) {
	user, err := database.Client.User.Query().Where(user.Username(request.Email)).Only(s.ctx);

	if err != nil {
		return "" ,errors.Error(ErrUserNotFound)
	}
	
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", errors.Error(ErrInvalidPassword)
	}

	// TODO: not everybody should be an admin!!!
	// TODO: pass user guid not email
	jwt, err := s.jwtService.GenerateToken(request.Email, role.Admin)

	if err != nil {
		return "", errors.Error(ErrInvalidClaims)
	}

	return jwt, nil;
}
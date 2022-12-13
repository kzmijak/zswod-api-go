package user

import (
	"context"

	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/user"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"golang.org/x/crypto/bcrypt"
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	ErrUserNotFound = "err_user_not_found: Could not match any user with provided username"
	ErrInvalidPassword = "err_invalid_password: Password for the user is incorrect"
)

func SignIn(ctx context.Context, request SignInRequest) (*ent.User, error) {
	user, err := database.Client.User.Query().Where(user.Username(request.Username)).Only(ctx);

	if err != nil {
		return nil ,errors.Error(ErrUserNotFound)
	}
	
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, errors.Error(ErrInvalidPassword)
	}

	return user, nil;
}
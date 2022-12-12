package user

import (
	"context"

	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

const (
	ErrCouldNotSaltPassword = "err_could_not_salt_password: Failed to encrypt the given password"
	ErrUserCreationFailed = "err_user_creation_failed: Failed to save the user in the database"
)

func CreateUser(ctx context.Context, request CreateUserRequest) (*ent.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Error(ErrCouldNotSaltPassword)
	}

	salt := string(hash)

	user, err := database.Client.User.Create().SetEmail(request.Email).SetPassword(salt).SetUsername(request.Username).Save(ctx)
	if err != nil {
		return nil, errors.Error(ErrUserCreationFailed)
	}

	return user, nil
}
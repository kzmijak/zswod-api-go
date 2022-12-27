package user

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/user"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

const (
	ErrCouldNotSaltPassword = "err_could_not_salt_password: Failed to encrypt the given password"
	ErrUserCreationFailed = "err_user_creation_failed: Failed to save the user in the database"
	ErrUserAlreadyExists = "err_user_already_exists: Failed to create user, user with provided email already exists"
)

func (s UserService) CreateUser(request CreateUserRequest) (*ent.User, error) {
	if alreadyExists, _ := database.Client.User.Query().Where(user.Email(request.Email)).Exist(s.ctx); alreadyExists == true {
		return nil, errors.Error(ErrUserAlreadyExists)
	}
	

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Error(ErrCouldNotSaltPassword)
	}

	salt := string(hash)

	user, err := database.Client.User.Create().
		SetID(uuid.New()).
		SetEmail(request.Email).
		SetPassword(salt).
		Save(s.ctx)
	if err != nil {
		return nil, errors.Error(ErrUserCreationFailed)
	}

	return user, nil
}
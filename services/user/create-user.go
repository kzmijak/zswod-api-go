package user

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/user"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
}

const (
	ErrCouldNotSaltPassword = "ErrCouldNotSaltPassword: Failed to encrypt the given password"
	ErrUserCreationFailed = "ErrUserCreationFailed: Failed to save the user in the database"
	ErrUserAlreadyExists = "ErrUserAlreadyExists: Failed to create user, user with provided email already exists"
	ErrInvalidRole = "ErrInvalidRole: Provided user does not exist"
)

func (s UserService) CreateUser(request CreateUserRequest) (*ent.User, error) {
	if _, exists := role.FromString(request.Role); !exists {
		return nil, errors.Error(ErrInvalidRole)
	}
	
	if alreadyExists, _ := database.Client.User.Query().Where(user.Email(request.Email)).Exist(s.ctx); alreadyExists {
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
		SetRolesID(request.Role).
		Save(s.ctx)
		
	if err != nil {
		return nil, errors.Error(ErrUserCreationFailed)
	}

	return user, nil
}
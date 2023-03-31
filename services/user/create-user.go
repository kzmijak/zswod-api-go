package user

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/user"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/utils/encryption"
)

const (
	ErrUserCreationFailed = "ErrUserCreationFailed: Failed to save the user in the database"
	ErrUserAlreadyExists = "ErrUserAlreadyExists: Failed to create user, user with provided email already exists"
)

type CreateUserPayload struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
}

func (s UserService) CreateUser(payload CreateUserPayload, tx *ent.Tx) (*ent.User, error) {
	if _, err := role.FromString(payload.Role); err != nil {
		return nil, err
	}

	existingUser, _ := s.GetUserByEmail(payload.Email, tx)
	if existingUser != nil {
		return nil, errors.Error(ErrUserAlreadyExists)
	}

	salt, err := encryption.HashString(payload.Password)
	if err != nil {
		return nil, err
	}

	user, err := tx.User.Create().
		SetID(uuid.New()).
		SetEmail(payload.Email).
		SetPassword(salt).
		SetRole(user.Role(payload.Role)).
		Save(s.ctx)
		
	if err != nil {
		return nil, errors.Error(ErrUserCreationFailed)
	}

	return user, nil
}
package userService

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/user"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrUserEmailQueryFailed = "ErrUserEmailQueryFailed: Failed to query user by email"
	ErrTokenCreationFailed = "ErrTokenCreationFailed: Failed to create token for this email"
)

func (s UserService) CreateResetPasswordToken(email string, tx *ent.Tx) (*ent.ResetPasswordToken, error) {

	owner, err := tx.User.Query().Where(user.Email(email)).First(s.ctx)
	if err != nil {
		return nil, errors.Error(ErrUserEmailQueryFailed)
	}

	token, err := tx.ResetPasswordToken.Create().
		SetCreatedAt(time.Now()).
		SetID(uuid.New()).
		SetOwner(owner).
		Save(s.ctx)
	if err != nil {
		return nil, errors.Error(ErrTokenCreationFailed)
	}

	return token, nil
}

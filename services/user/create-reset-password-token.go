package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/user"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrUserEmailQueryFailed = "ErrUserEmailQueryFailed: Failed to query user by email"
	ErrTokenCreationFailed = "ErrTokenCreationFailed: Failed to create token for this email"
)

func (s UserService) CreateResetPasswordToken(email string) (string, error) {

	owner, err := database.Client.User.Query().Where(user.Email(email)).First(s.ctx)
	if err != nil {
		return "", errors.Error(ErrUserEmailQueryFailed)
	}

	token, err := database.Client.ResetPasswordToken.Create().
		SetCreatedAt(time.Now()).
		SetID(uuid.New()).
		SetOwner(owner).
		Save(s.ctx)
	if err != nil {
		return "", errors.Error(ErrTokenCreationFailed)
	}

	return token.ID.String(), nil
}

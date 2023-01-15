package user

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/utils/encryption"
)

const (
	ErrCouldNotChangePassword = "ErrCouldNotChangePassword: Unable to change password"
)

type SetNewPasswordParams struct {
	UserId      uuid.UUID `json:"userId"`
	NewPassword string `json:"newPassword"`
}

func (s UserService) SetNewPassword(params SetNewPasswordParams) error {
	user, err := database.Client.User.Get(s.ctx, params.UserId)
	if err != nil {
		return errors.Error(ErrUserNotFound)
	}

	passwordHashed, err := encryption.HashString(params.NewPassword)
	if err != nil {
		return err
	}

	if err = user.Update().SetPassword(passwordHashed).Exec(s.ctx); err != nil {
		return errors.Error(ErrCouldNotChangePassword)
	}

	return nil
}
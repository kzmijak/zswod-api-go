package user

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/utils/encryption"
)

const (
	ErrCouldNotChangePassword = "ErrCouldNotChangePassword: Unable to change password"
)

type UpdateUserPasswordPayload struct {
	NewPassword string `json:"newPassword"`
}

func (s UserService) UpdateUserPassword(userId string, payload UpdateUserPasswordPayload, tx *ent.Tx) (*ent.User, error) {
	user, err := s.GetUser(userId, tx)
	if err != nil {
		return nil, err
	}

	passwordHashed, err := encryption.HashString(payload.NewPassword)
	if err != nil {
		return nil, err
	}

	userNew, err := user.Update().SetPassword(passwordHashed).Save(s.ctx)
	if err != nil {
		return nil, errors.Error(ErrCouldNotChangePassword)
	}

	return userNew, nil
}
package user

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotDeleteToken = "ErrCouldNotDeleteToken: Failed to remove token"
)

func (s UserService) InvalidateResetPasswordToken(token uuid.UUID) error {
	if err := database.Client.ResetPasswordToken.DeleteOneID(token).Exec(s.ctx); err != nil {
		return errors.Error(ErrCouldNotDeleteToken)
	}
	
	return nil
}
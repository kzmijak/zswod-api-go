package user

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotDeleteToken = "ErrCouldNotDeleteToken: Failed to remove token"
)

func (s UserService) InvalidateResetPasswordToken(token uuid.UUID, tx *ent.Tx) error {
	if err := tx.ResetPasswordToken.DeleteOneID(token).Exec(s.ctx); err != nil {
		return errors.Error(ErrCouldNotDeleteToken)
	}
	
	return nil
}
package user

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotDeleteToken = "ErrCouldNotDeleteToken: Failed to remove token"
)

func (s UserService) InvalidateResetPasswordToken(tokenId string, tx *ent.Tx) error {
	token, err := s.GetResetPasswordToken(tokenId, tx)
	if err != nil {
		return err
	}

	if err := tx.ResetPasswordToken.DeleteOne(token).Exec(s.ctx); err != nil {
		return errors.Error(ErrCouldNotDeleteToken)
	}
	
	return nil
}
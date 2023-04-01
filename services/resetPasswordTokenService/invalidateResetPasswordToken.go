package resetPasswordTokenService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotDeleteToken = "ErrCouldNotDeleteToken: Failed to remove token"
)

func (s ResetPasswordTokenService) InvalidateResetPasswordToken(tokenId uuid.UUID, tx *ent.Tx) error {
	token, err := s.tokenSelectors.SelectTokenById(tx, tokenId)
	if err != nil {
		return err
	}

	err = tx.ResetPasswordToken.DeleteOne(token).Exec(s.ctx);
	if err != nil {
		return errors.Error(ErrCouldNotDeleteToken)
	}
	
	return nil
}
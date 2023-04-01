package resetPasswordTokenService

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/resetPasswordTokenModel"
)

func (s ResetPasswordTokenService) GetResetPasswordTokenByEmail(ownerEmail string, tx *ent.Tx) (resetPasswordTokenModel.ResetPasswordTokenModel, error) {
	tokenEntity, err := s.tokenSelectors.SelectTokenByOwnerEmail(tx, ownerEmail)
	if err != nil {
		return resetPasswordTokenModel.Nil, err
	}

	return resetPasswordTokenModel.FromEntity(tokenEntity)
}
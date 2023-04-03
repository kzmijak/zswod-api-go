package resetPasswordTokenService

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/resetPasswordTokenModel"
)

func (s ResetPasswordTokenService) GetResetPasswordTokenByEmail(ownerEmail string, tx *ent.Tx) (resetPasswordTokenModel.ResetPasswordTokenModel, error) {
	return s.tokenSelectors.SelectTokenByOwnerEmail(tx, ownerEmail)
}
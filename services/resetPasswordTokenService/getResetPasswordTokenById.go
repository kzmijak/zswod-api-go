package resetPasswordTokenService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/resetPasswordTokenModel"
)

func (s ResetPasswordTokenService) GetResetPasswordTokenById(tokenId uuid.UUID, tx *ent.Tx) (resetPasswordTokenModel.ResetPasswordTokenModel, error) {
	return s.tokenSelectors.SelectTokenById(tx, tokenId)
}
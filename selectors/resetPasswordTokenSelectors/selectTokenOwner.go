package resetPasswordTokenSelectors

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/resetPasswordTokenQuery"
)

const (
	ErrCouldNotQueryForTokenOwner = "ErrCouldNotQueryForTokenOwner: Failed to query for reset password token owner"
)

func (s ResetPasswordTokenSelectors) SelectTokenOwner(tx *ent.Tx, tokenId uuid.UUID) (*ent.User, error) {
	ownerEntity, err := resetPasswordTokenQuery.FromTx(tx).
		QueryById(tokenId).
		QueryOwner().
		Only(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrCouldNotQueryForTokenOwner)
	}
	
	return ownerEntity, nil
}
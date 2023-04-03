package resetPasswordTokenSelectors

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/resetPasswordTokenModel"
	"github.com/kzmijak/zswod_api_go/query/resetPasswordTokenQuery"
)

const (
	ErrCouldNOtQueryTokenById = "ErrCouldNOtQueryTokenById: Failed to query token by ID. No token found." 
)

func (s ResetPasswordTokenSelectors) SelectTokenById(tx *ent.Tx, tokenId uuid.UUID) (resetPasswordTokenModel.ResetPasswordTokenModel, error) {
	tokenEntity, err := resetPasswordTokenQuery.FromTx(tx).
		QueryById(tokenId).
		QueryFullResetPasswordToken().
		Only(s.ctx)
		
	if err != nil {
		return resetPasswordTokenModel.Nil, err
	}

	return resetPasswordTokenModel.FromEntity(tokenEntity)
} 
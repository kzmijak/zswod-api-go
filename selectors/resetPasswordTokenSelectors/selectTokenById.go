package resetPasswordTokenSelectors

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/query/resetPasswordTokenQuery"
)

const (
	ErrCouldNOtQueryTokenById = "ErrCouldNOtQueryTokenById: Failed to query token by ID. No token found." 
)

func (s ResetPasswordTokenSelectors) SelectTokenById(tx *ent.Tx, tokenId uuid.UUID) (*ent.ResetPasswordToken, error) {
	tokenEntity, err := resetPasswordTokenQuery.FromTx(tx).
		QueryById(tokenId).
		QueryFullResetPasswordToken().
		Only(s.ctx)
		
	if err != nil {
		return nil, err
	}

	return tokenEntity, nil
} 
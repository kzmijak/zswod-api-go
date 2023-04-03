package resetPasswordTokenSelectors

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/resetPasswordTokenModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/resetPasswordTokenQuery"
	"github.com/kzmijak/zswod_api_go/query/userQuery"
)

const (
	ErrCouldNotQueryTokenByOwner = "ErrCouldNotQueryTokenByOwner: Failed to query for token owner"
)

func (s ResetPasswordTokenSelectors) SelectTokenByOwnerEmail(tx *ent.Tx, email string) (resetPasswordTokenModel.ResetPasswordTokenModel, error) {
	tokenEntQuery := userQuery.FromTx(tx).
		QueryByEmail(email).
		QueryResetPasswordToken()

	tokenEntity, err := resetPasswordTokenQuery.FromQuery(tokenEntQuery).
		QueryFullResetPasswordToken().
		Only(s.ctx)

	if err != nil {
		return resetPasswordTokenModel.Nil, errors.Error(ErrCouldNotQueryTokenByOwner)
	}

	return resetPasswordTokenModel.FromEntity(tokenEntity)
}
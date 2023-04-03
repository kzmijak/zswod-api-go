package userSelectors

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/userModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/userQuery"
)

const (
	ErrUserEmailQueryFail = "ErrUserEmailQueryFail: Failed to query for user. User does not exist."
)

func (s UserSelectors) SelectUserByEmail(tx *ent.Tx, email string) (userModel.UserModel, error) {
	userEntity, err := userQuery.FromTx(tx).
		QueryByEmail(email).
		QueryFullUser().
		Only(s.ctx)
	
	if err != nil {
		return userModel.Nil, errors.Error(ErrUserEmailQueryFail)
	}

	return userModel.FromEntity(userEntity)
}
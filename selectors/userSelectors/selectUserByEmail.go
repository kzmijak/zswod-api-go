package userSelectors

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/userQuery"
)

const (
	ErrUserEmailQueryFail = "ErrUserEmailQueryFail: Failed to query for user. User does not exist."
)

func (s UserSelectors) SelectUserByEmail(tx *ent.Tx, email string) (*ent.User, error) {
	userEntity, err := userQuery.FromTx(tx).
		QueryByEmail(email).
		QueryFullUser().
		Only(s.ctx)
	
	if err != nil {
		return nil, errors.Error(ErrUserEmailQueryFail)
	}

	return userEntity, nil
}
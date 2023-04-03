package userSelectors

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/userModel"
	"github.com/kzmijak/zswod_api_go/query/userQuery"
)

const (
	ErrCouldNotQueryUserById = "ErrCouldNotQueryUserById: User with given ID was not found" 
)

func (s UserSelectors) SelectUserById(tx *ent.Tx, userId uuid.UUID) (userModel.UserModel, error) {
	userEntity, err := userQuery.FromTx(tx).
		QueryById(userId).
		QueryFullUser().
		Only(s.ctx)
	
	if err != nil {
		return userModel.Nil, err
	}

	return userModel.FromEntity(userEntity)
}
package userService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/userModel"
)

func (s UserService) GetUserById(userId uuid.UUID, tx *ent.Tx) (userModel.UserModel, error) {
	return s.userSelectors.SelectUserById(tx, userId)
}
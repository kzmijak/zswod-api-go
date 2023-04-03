package userModel

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/imageModel"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/pickers/userEntityPicker"
)

type UserModel struct {
	ID   uuid.UUID `json:"id"`
	Email string `json:"email"`
	Role role.Role `json:"role"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
  Avatar imageModel.Image `json:"avatar"`
}

var Nil = UserModel{}

func FromEntity(userEntity *ent.User) (UserModel, error) {
	userRole, err := role.FromString(userEntity.Role.String())
	if err != nil {
		return Nil, err
	}

	var userAvatarModel imageModel.Image
	userAvatarEntity, _ := userEntityPicker.PickAvatarImageEntity(userEntity)
	if userAvatarEntity != nil {
		userAvatarModel, err = imageModel.FromEntity(userAvatarEntity)
		if err != nil {
			return Nil, err
		}
	}


	return UserModel{
		ID: userEntity.ID,
		Email: userEntity.Email,
		Role: userRole,
		FirstName: userEntity.FirstName,
		LastName: userEntity.LastName,
		Avatar: userAvatarModel,
	}, nil
}
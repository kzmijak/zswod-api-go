package resetPasswordTokenModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/userModel"
	"github.com/kzmijak/zswod_api_go/pickers/resetPasswordTokenEntityPicker"
)

type ResetPasswordTokenModel struct {
	ID uuid.UUID `json:"id"`
	CreateTime time.Time `json:"createdAt"`
	ExpiryTime time.Time `json:"expiryTime"`
	Owner userModel.UserModel `json:"owner"`
}

var Nil = ResetPasswordTokenModel{}

func FromEntity(tokenEntity *ent.ResetPasswordToken) (ResetPasswordTokenModel, error) {
	ownerEntity, err := resetPasswordTokenEntityPicker.PickOwnerUserEntity(tokenEntity) 
	if err != nil {
		return Nil, err
	}
	ownerModel, err := userModel.FromEntity(ownerEntity)
	if err != nil {
		return Nil, err
	}

	return ResetPasswordTokenModel{
		ID: tokenEntity.ID,
		CreateTime: tokenEntity.CreateTime,
		ExpiryTime: tokenEntity.ExpiryTime,
		Owner: ownerModel,
		}, nil
}
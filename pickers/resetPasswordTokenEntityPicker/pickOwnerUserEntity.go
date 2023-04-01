package resetPasswordTokenEntityPicker

import "github.com/kzmijak/zswod_api_go/ent"

const (
	ErrCouldNotQueryOwner = "ErrCouldNotQueryOwner: Could not query for token owner"
)

func PickOwnerUserEntity(entity *ent.ResetPasswordToken) (*ent.User, error) {
	userEntity, err := entity.Edges.OwnerOrErr()
	if err != nil {
		return nil, err
	}

	return userEntity, nil
}
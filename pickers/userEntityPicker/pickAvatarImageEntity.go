package userEntityPicker

import "github.com/kzmijak/zswod_api_go/ent"

func PickAvatarImageEntity(userEntity *ent.User) (*ent.Image, error) {
	imageEntity, err := userEntity.Edges.AvatarOrErr()
	if err != nil {
		return nil, err
	}

	return imageEntity, nil
}
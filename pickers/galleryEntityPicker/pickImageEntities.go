package galleryEntityPicker

import "github.com/kzmijak/zswod_api_go/ent"

const (
	ErrCouldNotPickImagesFromGallery = "ErrCouldNotPickImagesFromGallery: Failed to pick images from gallery"
)

func PickImageEntities(galleryEntity *ent.Gallery) ([]*ent.Image, error) {
	imageEntities, err := galleryEntity.Edges.ImagesOrErr()
	if err != nil {
		return nil, err
	}

	return imageEntities, nil
}
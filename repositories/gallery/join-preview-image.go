package galleryRepo

import (
	"github.com/kzmijak/zswod_api_go/ent"
	imageRepo "github.com/kzmijak/zswod_api_go/repositories/image"
)

func (query GalleryQuery) JoinPreviewImage() GalleryQuery {
	query.WithImages(func(iq *ent.ImageQuery) {
		imageRepo.QueryImage(iq).QueryOrderedWithBlobId().Limit(1)
	})

	return query
}
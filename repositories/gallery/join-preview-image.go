package galleryRepo

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/image"
	imageRepo "github.com/kzmijak/zswod_api_go/repositories/image"
)

func (query GalleryQuery) JoinPreviewImage() GalleryQuery {
	query.WithImages(func(iq *ent.ImageQuery) {
		imageRepo.QueryImage(iq).QueryOrderedWithBlobId().Where(image.IsPreview(true))
	})

	return query
}
package galleryQuery

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/image"
	"github.com/kzmijak/zswod_api_go/query/imageQuery"
)

func (query GalleryQuery) JoinPreviewImage() GalleryQuery {
	query.WithImages(func(iq *ent.ImageQuery) {
		imageQuery.FromQuery(iq).QueryOrderedWithBlobId().Where(image.Order(0))
	})

	return query
}
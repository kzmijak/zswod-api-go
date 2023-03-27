package galleryQuery

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/query/imageQuery"
)

func (query GalleryQuery) JoinAllImagesToGallery() GalleryQuery {
	query.WithImages(func(iq *ent.ImageQuery) {
		imageQuery.FromQuery(iq).QueryOrdered()
	})
	
	return query
}
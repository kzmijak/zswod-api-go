package galleryQuery

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/gallery"
)

func (query GalleryQuery) QueryById(galleryId uuid.UUID) GalleryQuery {
	galleryQuery := query.Where(gallery.ID(galleryId))
	return FromQuery(galleryQuery)
}
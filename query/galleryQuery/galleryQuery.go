package galleryQuery

import "github.com/kzmijak/zswod_api_go/ent"

type GalleryQuery struct {
	*ent.GalleryQuery
}

func FromQuery (query *ent.GalleryQuery) GalleryQuery {
	return GalleryQuery{
		query,
	}
}

func FromTx(tx *ent.Tx) GalleryQuery {
	return FromQuery(tx.Gallery.Query())
}
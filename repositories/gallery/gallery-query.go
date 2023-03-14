package galleryRepo

import "github.com/kzmijak/zswod_api_go/ent"

type GalleryQuery struct {
	*ent.GalleryQuery
}

func QueryGallery (query *ent.GalleryQuery) GalleryQuery {
	return GalleryQuery{
		query,
	}
}

func GalleryTx(tx *ent.Tx) GalleryQuery {
	return QueryGallery(tx.Gallery.Query())
}
package blobQuery

import "github.com/kzmijak/zswod_api_go/ent/blob"

func (bq BlobQuery) QueryPictures() BlobQuery {
	query := bq.Where(blob.TypeEQ(blob.TypePicture))
	return FromQuery(query)
}
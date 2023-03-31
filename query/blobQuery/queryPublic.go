package blobQuery

import "github.com/kzmijak/zswod_api_go/ent/blob"

func (bq BlobQuery) QueryPublic() BlobQuery {
	query := bq.Where(blob.IsPublic(true))
	return FromQuery(query)
}

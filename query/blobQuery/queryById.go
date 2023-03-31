package blobQuery

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/blob"
)

func (bq BlobQuery) QueryById(id uuid.UUID) BlobQuery {
	blobQuery := bq.Where(blob.ID(id))
	return FromQuery(blobQuery)
}

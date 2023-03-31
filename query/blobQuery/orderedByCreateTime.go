package blobQuery

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/blob"
)

func (bq BlobQuery) OrderedByCreateTime() BlobQuery {
	query := bq.Order(ent.Desc(blob.FieldCreateTime))
	return FromQuery(query)
}
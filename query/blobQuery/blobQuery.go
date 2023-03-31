package blobQuery

import "github.com/kzmijak/zswod_api_go/ent"

type BlobQuery struct {
	*ent.BlobQuery
}

func FromQuery (query *ent.BlobQuery) BlobQuery {
	bq := BlobQuery{
		query,
	}
	return bq;
}

func FromTx(tx *ent.Tx) BlobQuery {
	return FromQuery(tx.Blob.Query())
}
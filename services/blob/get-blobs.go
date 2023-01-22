package blob

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotQuery = "ErrCouldNotQuery: Failed to query for blobs"
)

type GetBlobsReturnType struct {
	Blobs []*ent.Blob `json:"blobs"`
	Eof bool `json:"eof"` 
}

func (s BlobService) GetBlobs(amount int, offset int) (*GetBlobsReturnType, error) {
	queryBase := database.Client.Blob.Query()
	
	count, err := queryBase.Count(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrCouldNotQuery)
	}

	blobs, err := queryBase.
		Limit(amount).
		Offset(offset).
		Order(ent.Desc(blob.FieldCreatedAt)).
		All(s.ctx)
	
	if err != nil {
		return nil, errors.Error(ErrCouldNotQuery)
	}

	eof := amount + offset >= count


	return &GetBlobsReturnType{
		Blobs: blobs,
		Eof: eof,
	}, nil
}
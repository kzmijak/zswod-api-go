package blobService

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/blobModel"
)

type GetBlobsReturnType struct {
	Blobs []blobModel.BlobModel `json:"blobs"`
	Eof bool `json:"eof"` 
}

var Nil = GetBlobsReturnType{}

func (s BlobService) GetBlobs(amount int, offset int, tx *ent.Tx) (GetBlobsReturnType, error) {
	blobEntities, eof, err := s.selectors.SelectPublicPictureBlobs(tx, offset, amount)
	if err != nil {
		return Nil, err
	}

	blobModels, err := blobModel.ArrayFromEntities(blobEntities)
	if err != nil {
		return Nil, err
	}

	return GetBlobsReturnType{
		Blobs: blobModels,
		Eof: eof,
	}, nil
}
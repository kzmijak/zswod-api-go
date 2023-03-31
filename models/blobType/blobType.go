package blobType

import (
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/utils/array"
)

const (
	ErrBlobTypeDoesNotExist = "ErrBlobTypeDoesNotExist: Blob type does not exist"
)

type BlobType string

const (
	Unknown    BlobType = "Unknown"
	Attachment BlobType = "Attachment"
	Picture    BlobType = "Picture"
)

var BlobTypes = []BlobType{ Attachment, Picture }

func (blobType BlobType) String() string {
	return string(blobType)
}

func FromString(blobTypeString string) (BlobType, error) {
	blobType := BlobType(blobTypeString)
	exists := array.Includes(BlobTypes, blobType)
	if !exists {
		return Unknown, errors.Error(ErrBlobTypeDoesNotExist)
	}

	return blobType, nil
}

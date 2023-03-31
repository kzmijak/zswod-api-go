package blobType

import (
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/utils/bimap"
)

const (
	ErrBlobTypeDoesNotExist = "ErrBlobTypeDoesNotExist: Blob type does not exist"
)

type BlobType int

const (
	Unknown    BlobType = iota - 1
	Attachment BlobType = iota - 1
	Picture
)

var blobTypeStringMap = bimap.NewBiMap[BlobType]().
	WithMember(Unknown, "Unknown").
	WithMember(Attachment, "Attachment").
	WithMember(Picture, "Picture")

func (blobType BlobType) String() string {
	return blobTypeStringMap.GetByKey(blobType)
}

func FromString(blobTypeString string) (BlobType, error) {
	blobType, exists := blobTypeStringMap.GetByValue(blobTypeString)
	if !exists {
		return Unknown, errors.Error(ErrBlobTypeDoesNotExist)
	}

	return blobType, nil
}

func FromId(blobTypeId int) (BlobType, error) {
	exists := blobTypeStringMap.GetLength() > blobTypeId
	if !exists {
		return Unknown, errors.Error(ErrBlobTypeDoesNotExist)
	}

	return BlobType(blobTypeId), nil
}
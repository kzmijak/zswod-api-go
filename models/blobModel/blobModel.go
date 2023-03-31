package blobModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/blobType"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type BlobModel struct {
	ID          uuid.UUID  `json:"id"`
	CreateTime  time.Time  `json:"createTime"`
	Blob        []byte     `json:"blob"`
	ContentType string     `json:"contentType"`
	IsPublic 	  bool 			 `json:"isPublic"`
	Type blobType.BlobType `json:"blobType"`
}

var Nil = BlobModel{}

func FromEntity(blobEntity *ent.Blob) (BlobModel, error) {
	blobType, err := blobType.FromString(blobEntity.Type.String())
	if err != nil {
		return Nil, err
	}

	return BlobModel{
		ID: blobEntity.ID,
		CreateTime: blobEntity.CreateTime,
		Blob: blobEntity.Blob,
		ContentType: blobEntity.ContentType,
		IsPublic: blobEntity.IsPublic,
		Type: blobType,
	}, nil
}

var ArrayFromEntities = arraymap.CreateArrayMapper(FromEntity)

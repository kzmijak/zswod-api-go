package blobModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/models/blobType"
)

type BlobModel struct {
	ID          uuid.UUID  `json:"id"`
	CreateTime  time.Time  `json:"createTime"`
	Blob        []byte     `json:"blob"`
	ContentType string     `json:"contentType"`
	IsPublic 	  bool 			 `json:"isPublic"`
	Type blobType.BlobType `json:"blobType"`
}
package attachmentEntityPicker

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotPickBlobFromAttachment = "ErrCouldNotPickBlobFromAttachment: Failed to pick blob from attachment entity" 
)

func PickBlobEntity(attachmentEntity *ent.Attachment) (*ent.Blob, error) {
	blob, err := attachmentEntity.Edges.BlobOrErr()
	if err != nil {
		return nil, errors.Error(ErrCouldNotPickBlobFromAttachment)
	}

	return blob, nil
}
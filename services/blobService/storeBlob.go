package blobService

import (
	"io"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/models/blobModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrFileTooLarge = "ErrFileTooLarge: File too large, max weight: 4MB"
	ErrFileOpenFailed = "ErrFileOpenFailed: Failed to open file"
	ErrFileReadFailed = "ErrFileReadFailed: Failed to read file"
	ErrCouldNotSaveToDb = "ErrCouldNotSaveToDb: Failed to save to database"
)

func (s BlobService) StoreBlob(file *multipart.FileHeader, tx *ent.Tx) (blobModel.BlobModel, error)  {
	const FOUR_MEGABYTES = 4 << 20
	if file.Size > FOUR_MEGABYTES {
		return blobModel.Nil, errors.Error(ErrFileTooLarge)
	}
	
	content, err := file.Open()
	
	contentType := file.Header.Get("Content-Type")
	if err != nil {
		return blobModel.Nil, errors.Error(ErrFileOpenFailed)
	}

	byteContainer, err := io.ReadAll(content)
	if err != nil {
		return blobModel.Nil, errors.Error(ErrFileReadFailed)
	}

	newBlobEntity, err := tx.Blob.
		Create().
		SetID(uuid.New()).
		SetIsPublic(true).
		SetType(blob.TypePicture).
		SetBlob(byteContainer).
		SetTitle(file.Filename).
		SetContentType(contentType).
		SetCreateTime(time.Now()).
		Save(s.ctx)

	if err != nil {
		return blobModel.Nil, errors.Error(ErrCouldNotSaveToDb)
	}

	return blobModel.FromEntity(newBlobEntity)
}
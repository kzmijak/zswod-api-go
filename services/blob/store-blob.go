package blob

import (
	"io/ioutil"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrFileTooLarge = "ErrFileTooLarge: File too large, max weight: 4MB"
	ErrFileOpenFailed = "ErrFileOpenFailed: Failed to open file"
	ErrFileReadFailed = "ErrFileReadFailed: Failed to read file"
	ErrCouldNotSaveToDb = "ErrCouldNotSaveToDb: Failed to save to database"
)

func (s BlobService) StoreBlob(file *multipart.FileHeader) (*ent.Blob, error) {
	const FOUR_MEGABYTES = 4 << 20
	if file.Size > FOUR_MEGABYTES {
		return nil, errors.Error(ErrFileTooLarge)
	}
	
	content, err := file.Open()
	contentType := file.Header.Get("Content-Type")
	if err != nil {
		return nil, errors.Error(ErrFileOpenFailed)
	}

	byteContainer, err := ioutil.ReadAll(content)
	if err != nil {
		return nil, errors.Error(ErrFileReadFailed)
	}

	response, err := database.Client.Blob.
		Create().
		SetID(uuid.New()).
		SetBlob(byteContainer).
		SetContentType(contentType).
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrCouldNotSaveToDb)
	}

	return response, nil
}
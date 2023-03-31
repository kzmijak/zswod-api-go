package selectors

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/blobQuery"
)

const (
	ErrPublicPicturesQueryFail = "ErrPublicPicturesQueryFail: Failed to query for public picture blobs"
)

func (s BlobSelector) SelectPublicPictureBlobs(tx *ent.Tx, offset int, amount int) (blobEntities []*ent.Blob, eof bool, err error) {
	query := blobQuery.FromTx(tx).
		QueryPublic().
		QueryPictures().
		OrderedByCreateTime().
		Clone()
		
	count, err := query.Count(s.ctx)
	if err != nil {
		err = errors.Error(ErrPublicPicturesQueryFail)
		return
	}
	
	blobEntities, err = query.Limit(amount).
		Offset(offset).
		All(s.ctx)

	if err != nil {
		err = errors.Error(ErrPublicPicturesQueryFail)
		return
	}

	eof = amount + offset >= count


	return
}
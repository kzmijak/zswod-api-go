package articleHeaderModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/imageModel"
	"github.com/kzmijak/zswod_api_go/pickers/articleEntityPicker"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type ArticleHeaderModel struct {
	ID              uuid.UUID        `json:"id,omitempty"`
	Title           string           `json:"title,omitempty"`
	Short           string           `json:"short,omitempty"`
	UploadDate      time.Time        `json:"uploadDate,omitempty"`
	TitleNormalized string           `json:"titleNormalized,omitempty"`
	PreviewImage    imageModel.Image `json:"previewImage"`
}

var Nil = ArticleHeaderModel{}

func FromArticleEntity(articleEntity *ent.Article) (ArticleHeaderModel, error) {
	previewImageEntity, err := articleEntityPicker.PickGalleryPreviewImage(articleEntity)
	if err != nil {
		return Nil, err
	}

	previewImageModel, err := imageModel.FromEntity(previewImageEntity)
	if err != nil {
		return Nil, err
	}

	return ArticleHeaderModel{
		ID: articleEntity.ID,
		Title: articleEntity.Title,
		Short: articleEntity.Short,
		UploadDate: articleEntity.UploadDate,
		TitleNormalized: articleEntity.TitleNormalized,
		PreviewImage: previewImageModel,
	}, nil
}

var ArrayFromEntities = arraymap.CreateArrayMapper(FromArticleEntity)
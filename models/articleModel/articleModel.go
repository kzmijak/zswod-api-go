package articleModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/galleryModel"
	"github.com/kzmijak/zswod_api_go/pickers/articleEntityPicker"
)

type ArticleModel struct {
	ID         uuid.UUID                 `json:"id,omitempty"`
	CreateTime time.Time                 `json:"createTime,omitempty"`
	UpdateTime time.Time                 `json:"updateTime,omitempty"`
	Title      string                    `json:"title,omitempty"`
	Short      string                    `json:"short,omitempty"`
	Content    string                    `json:"content,omitempty"`
	Gallery    galleryModel.GalleryModel `json:"gallery"`
}

var Nil = ArticleModel{}

func FromEntity(articleEntity *ent.Article) (ArticleModel, error) {
	galleryEntity, err := articleEntityPicker.PickGalleryEntity(articleEntity)
	if err != nil {
		return Nil, err
	}

	galleryModel, err := galleryModel.FromEntity(galleryEntity)
	if err != nil {
		return Nil, err
	}

	return ArticleModel{
		ID:         articleEntity.ID,
		Title:      articleEntity.Title,
		Short:      articleEntity.Short,
		Content:    articleEntity.Content,
		CreateTime: articleEntity.CreateTime,
		UpdateTime: articleEntity.UpdateTime,
		Gallery:    galleryModel,
	}, nil
}
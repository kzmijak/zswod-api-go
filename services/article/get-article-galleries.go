package article

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/article"
	entImage "github.com/kzmijak/zswod_api_go/ent/image"
	"github.com/kzmijak/zswod_api_go/models/image"
	"github.com/kzmijak/zswod_api_go/modules/database"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type ArticleGallery struct {
	ID              uuid.UUID     `json:"id,omitempty"`
	Title           string        `json:"title,omitempty"`
	UploadDate      time.Time     `json:"uploadDate,omitempty"`
	TitleNormalized string        `json:"titleNormalized,omitempty"`
	Images 					[]image.Image `json:"images"`
}

func (s ArticleService) GetArticleGalleries(amount int, offset int) ([]ArticleGallery, error) {
	entities, err := database.Client.Article.Query().
		Order(ent.Desc(article.FieldUploadDate)).
		Limit(amount).
		Offset(offset).
		WithTitleNormalized().
		WithImages(func(iq *ent.ImageQuery) {
			iq.Order(ent.Asc(entImage.FieldOrder)).All(s.ctx)
		}).
		All(s.ctx)

	if err != nil {
		return nil, err
	}

	return arraymap.CreateArrayMapper(mapEntityToResponse)(entities), nil
}

func mapEntityToResponse(e *ent.Article) ArticleGallery {
		return ArticleGallery{
			ID: e.ID,
			Title: e.Title,
			UploadDate: e.UploadDate,
			TitleNormalized: e.Edges.TitleNormalized.TitleNormalized,
			Images: image.ArrayFromEntities(e.Edges.Images),
		}
	}
package article

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/samber/lo"
)

const (
	ErrQueryFail = "ErrQueryFail: Failed to execute query"
)

type ArticleHeader struct  {
	ID uuid.UUID `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Short string `json:"short,omitempty"`
	UploadDate time.Time `json:"uploadDate,omitempty"`
	TitleNormalized string `json:"titleNormalized,omitempty"`
}
func (s ArticleService) GetArticleHeaders(amount int, offset int) ([]ArticleHeader, error) {
	articles, err := database.Client.Article.Query().
		Order(ent.Desc(article.FieldUploadDate)).
		Limit(amount).
		Offset(offset).
		WithTitleNormalized().
		All(s.ctx)

	articleHeaders := lo.Map(articles, func(a *ent.Article, _ int) ArticleHeader {
		return ArticleHeader{
			ID: a.ID,
			Title: a.Title,
			Short: a.Short,
			UploadDate: a.UploadDate,
			TitleNormalized: a.Edges.TitleNormalized.TitleNormalized,
		}
	})

	if err != nil {
		return nil, errors.Error(ErrQueryFail)
	}

	return articleHeaders, nil
}
package article

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/article"
	imageMdl "github.com/kzmijak/zswod_api_go/models/image"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	articleRepo "github.com/kzmijak/zswod_api_go/repositories/article"
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
	Images []imageMdl.Image `json:"images"`
}

func (s ArticleService) GetArticleHeaders(amount int, offset int, tx *ent.Tx) ([]ArticleHeader, error) {
	articles, err := articleRepo.ArticleTx(tx).
		JoinAllImagesToArticle().
		Order(ent.Desc(article.FieldUploadDate)).
		Limit(amount).
		Offset(offset).
		All(s.ctx)

	articleHeaders := lo.Map(articles, func(a *ent.Article, _ int) ArticleHeader {

		return ArticleHeader{
			ID: a.ID,
			Title: a.Title,
			Short: a.Short,
			UploadDate: a.UploadDate,
			TitleNormalized: a.TitleNormalized,
			Images: imageMdl.ArrayFromEntities(a.Edges.Gallery.Edges.Images),
		}
	})

	if err != nil {
		return nil, errors.Error(ErrQueryFail)
	}

	return articleHeaders, nil
}
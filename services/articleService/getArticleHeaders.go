package articleService

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/models/articleHeaderModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	articleRepo "github.com/kzmijak/zswod_api_go/repositories/article"
)

const (
	ErrFailedToQueryArticleHeaders = "ErrFailedToQueryArticleHeaders: Failed to query for article headers"
)

func (s ArticleService) GetArticleHeaders(amount int, offset int, tx *ent.Tx) ([]articleHeaderModel.ArticleHeaderModel, error) {
	articles, err := articleRepo.ArticleTx(tx).
		JoinAllImagesToArticle().
		Order(ent.Desc(article.FieldUploadDate)).
		Limit(amount).
		Offset(offset).
		All(s.ctx)

	if err != nil { 
		return nil, errors.Error(ErrFailedToQueryArticleHeaders)
	}

	return articleHeaderModel.ArrayFromEntities(articles)
}
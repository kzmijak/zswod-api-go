package articleService

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/articleHeaderModel"
)

const (
	ErrFailedToQueryArticleHeaders = "ErrFailedToQueryArticleHeaders: Failed to query for article headers"
)

func (s ArticleService) GetArticleHeaders(amount int, offset int, tx *ent.Tx) ([]articleHeaderModel.ArticleHeaderModel, error) {
	articleEntities, err := s.selectors.SelectPublishedArticles(tx, offset, amount)
	if err != nil {
		return nil, err
	}

	return articleHeaderModel.ArrayFromEntities(articleEntities)
}
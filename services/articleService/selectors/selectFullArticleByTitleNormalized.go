package selectors

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/articleQuery"
)

const (
	ErrArticleTitleNormalizedNotFound = "ErrArticleTitleNormalizedNotFound: Could not find article by given Title (normalized)"
)

func (s ArticleSelector) SelectFullArticleByTitleNormalized(tx *ent.Tx, titleNormalized string) (*ent.Article, error) {
	articleEntity, err := articleQuery.FromTx(tx).
		JoinAllImagesToArticle().
		QueryByTitleNormalized(titleNormalized).
		Only(s.ctx)
		
	if err != nil {
		return nil, errors.Error(ErrArticleIdNotFound)
	}

	return articleEntity, nil
}
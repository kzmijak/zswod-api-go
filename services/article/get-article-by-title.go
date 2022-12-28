package article

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/articletitleguid"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrArticleNotFound = "ErrArticleNotFound: Article with given name not found"
	ErrArticleNotMapped = "ErrArticleNotMapped: Name lacks connection with article"
)

func (s ArticleService) GetArticleByTitle(title string) (*ent.Article, error) {
	articleTitleGuid, err := database.Client.ArticleTitleGuid.Query().Where(articletitleguid.TitleNormalized(title)).Only(s.ctx)
	
	if err != nil {
		return nil, errors.Error(ErrArticleNotFound)
	}

	article, err := articleTitleGuid.QueryArticle().Only(s.ctx)
	if err != nil {
		return nil, errors.Error(ErrArticleNotMapped)
	}

	return article, nil
}
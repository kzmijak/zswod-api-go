package article

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/article"
	imageMdl "github.com/kzmijak/zswod_api_go/models/image"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrArticleTitleNotFound = "ErrArticleTitleNotFound: Could not find article by title"
)

type GetArticleByTitleResponse struct {
	Article article.Article `json:"article"`
	Images []imageMdl.Image `json:"images"`
}
func (s ArticleService) GetArticleByTitle(titleNormalized string, tx *ent.Tx) (*GetArticleByTitleResponse, error) {
	articleEntity, err := s.getArticleEntityByTitle(titleNormalized, tx)

	if err != nil {
		return nil, errors.Error(ErrArticleTitleNotFound)
	}

	imageModels := imageMdl.ArrayFromEntities(articleEntity.Edges.Images)


	return &GetArticleByTitleResponse{
		Article: article.FromEntity(*articleEntity),
		Images: imageModels,
	}, nil
}
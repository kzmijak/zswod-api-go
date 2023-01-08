package article

import (
	"github.com/kzmijak/zswod_api_go/models/article"
	imageMdl "github.com/kzmijak/zswod_api_go/models/image"
)

const (
	ErrArticleNotMapped = "ErrArticleNotMapped: Name lacks connection with article"
)

type GetArticleByTitleResponse struct {
	Article article.Article `json:"article"`
	Images []imageMdl.Image `json:"images"`
}
func (s ArticleService) GetArticleByTitle(titleNormalized string) (*GetArticleByTitleResponse, error) {
	articleEntity, err := s.getArticleEntityByTitle(titleNormalized)

	if err != nil {
		return nil, err
	}

	imageModels := imageMdl.ArrayFromEntities(articleEntity.Edges.Images)


	return &GetArticleByTitleResponse{
		Article: article.FromEntity(*articleEntity),
		Images: imageModels,
	}, nil
}
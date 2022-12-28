package article

import "context"

type ArticleService struct {
	ctx context.Context
}

func New() ArticleService {
	return ArticleService{}
}

func (s ArticleService) WithContext(ctx context.Context) ArticleService {
	s.ctx = ctx
	return s
}
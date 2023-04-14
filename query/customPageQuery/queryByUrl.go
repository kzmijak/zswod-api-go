package customPageQuery

import "github.com/kzmijak/zswod_api_go/ent/custompage"

func (query CustomPageQuery) QueryByUrl(url string) CustomPageQuery {
	customPageQuery := query.Where(custompage.TitleNormalized(url))
	return FromQuery(customPageQuery)
}
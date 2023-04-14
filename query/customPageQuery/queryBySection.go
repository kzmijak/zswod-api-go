package customPageQuery

import "github.com/kzmijak/zswod_api_go/ent/custompage"

func (query CustomPageQuery) QueryBySection(section string) CustomPageQuery {
	customPageQuery := query.Where(custompage.Section(section))
	return FromQuery(customPageQuery)
}
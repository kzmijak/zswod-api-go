package customPageHeaderModel

import (
	"github.com/kzmijak/zswod_api_go/ent"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)


type CustomPageHeaderModel struct {
	Url string `json:"url"`
	Section 				string `json:"section"`
	Title           string `json:"title"`
	IconId          string `json:"icon"`
	IsExternal      bool   `json:"isExternal"`
	Link            string `json:"link,omitempty"`
}

var Nil = CustomPageHeaderModel{}

func FromCustomPageEntity(customPageEntity *ent.CustomPage) (CustomPageHeaderModel, error) {
	return CustomPageHeaderModel{
		Url: customPageEntity.URL,
		Title: customPageEntity.Title,
		IconId: customPageEntity.IconId,
		IsExternal: customPageEntity.IsExternal,
		Link: customPageEntity.Link,
		Section: customPageEntity.Section,
	}, nil
}

var ArrayFromEntities = arraymap.CreateArrayMapper(FromCustomPageEntity)
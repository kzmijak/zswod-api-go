package customPageHeaderModel

import (
	"github.com/kzmijak/zswod_api_go/ent"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type CustomPageHeaderItem struct {
	TitleNormalized string `json:"titleNormalized"`
	Section 				string `json:"section"`
	Title           string `json:"title"`
	IconId          string `json:"icon"`
	IsExternal      bool   `json:"isExternal"`
	Link            string `json:"link,omitempty"`
}

type CustomPageHeaderModel struct {
	TitleNormalized string `json:"titleNormalized"`
	Section 				string `json:"section"`
	Title           string `json:"title"`
	IconId          string `json:"icon"`
	IsExternal      bool   `json:"isExternal"`
	Link            string `json:"link,omitempty"`
}

var Nil = CustomPageHeaderModel{}

func FromCustomPageEntity(customPageEntity *ent.CustomPage) (CustomPageHeaderModel, error) {
	return CustomPageHeaderModel{
		TitleNormalized: customPageEntity.TitleNormalized,
		Title: customPageEntity.Title,
		IconId: customPageEntity.IconId,
		IsExternal: customPageEntity.IsExternal,
		Link: customPageEntity.Link,
		Section: customPageEntity.Section,
	}, nil
}

var ArrayFromEntities = arraymap.CreateArrayMapper(FromCustomPageEntity)
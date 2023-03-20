package customPageHeaderModel

type CustomPageHeaderModel struct {
	TitleNormalized string `json:"titleNormalized"`
	Title           string `json:"title"`
	IconId          string `json:"icon"`
	IsExternal      bool   `json:"isExternal"`
	Link            string `json:"link,omitempty"`
}
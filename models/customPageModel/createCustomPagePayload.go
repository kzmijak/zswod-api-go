package customPageModel

type CreateCustomPagePayload struct {
	Section string `json:"section"`
	Title   string `json:"title"`
	Content string `json:"content"`
	IconId  string `json:"iconId"`
}

func NewCreateCustomPagePayload() CreateCustomPagePayload {
	return CreateCustomPagePayload{}
}

func (model CreateCustomPagePayload) WithSection(section string) CreateCustomPagePayload {
	model.Section = section
	return model
}
func (model CreateCustomPagePayload) WithTitle(title string) CreateCustomPagePayload {
	model.Title = title
	return model
}
func (model CreateCustomPagePayload) WithContent(content string) CreateCustomPagePayload {
	model.Content = content
	return model
}
func (model CreateCustomPagePayload) WithIconId(iconId string) CreateCustomPagePayload {
	model.IconId = iconId
	return model
}

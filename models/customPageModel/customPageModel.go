package customPageModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/attachmentModel"
)

type CustomPageModel struct {
	ID uuid.UUID `json:"id"`
	UpdateTime  time.Time                         `json:"updateTime"`
	Section 	  string  `json:"section"`
	Title       string                            `json:"title"`
	Content     string                            `json:"content"`
	Attachments []attachmentModel.AttachmentModel `json:"attachments"`
	Url string `json:"url"`
}

var Nil = CustomPageModel{}

func FromCustomPageEntity(customPageEntity *ent.CustomPage) (CustomPageModel, error) {
	attachmentEntities := customPageEntity.Edges.Attachments
	attachmentModels, err := attachmentModel.ArrayFromEntities(attachmentEntities)

	if err != nil {
		return Nil, err
	}

	return CustomPageModel{
		ID: customPageEntity.ID,
		UpdateTime: customPageEntity.UpdateTime,
		Section: customPageEntity.Section,
		Title: customPageEntity.Title,
		Content: customPageEntity.Content,
		Attachments: attachmentModels,
		Url: customPageEntity.URL,
	}, nil
}

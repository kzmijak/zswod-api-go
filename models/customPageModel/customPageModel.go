package customPageModel

import (
	"time"

	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/attachmentModel"
)

type CustomPageModel struct {
	UpdateTime  time.Time                         `json:"updateTime"`
	Section 	  string  `json:"section"`
	Title       string                            `json:"title"`
	Content     string                            `json:"content"`
	Attachments []attachmentModel.AttachmentModel `json:"attachments"`
	TitleNormalized string `json:"titleNormalized"`
}

var Nil = CustomPageModel{}

func FromCustomPageEntity(customPageEntity *ent.CustomPage) (CustomPageModel, error) {
	attachmentEntites := customPageEntity.Edges.Attachments
	attachmentModels, err := attachmentModel.ArrayFromEntities(attachmentEntites)

	if err != nil {
		return Nil, err
	}

	return CustomPageModel{
		UpdateTime: customPageEntity.UpdateTime,
		Section: customPageEntity.Section,
		Title: customPageEntity.Title,
		Content: customPageEntity.Content,
		Attachments: attachmentModels,
		TitleNormalized: customPageEntity.TitleNormalized,
	}, nil
}

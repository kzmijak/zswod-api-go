package customPageModel

import (
	"time"

	"github.com/kzmijak/zswod_api_go/models/attachmentModel"
)

type CustomPageModel struct {
	UpdateTime  time.Time                         `json:"updateTime"`
	Title       string                            `json:"title"`
	Content     string                            `json:"content"`
	Attachments []attachmentModel.AttachmentModel `json:"attachments"`
}
package userController

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/mailer"
)

type ResetPasswordBody struct {
	Email string `json:"email"`
}

func (c UserController) ResetPassword(ctx *gin.Context) {
	var requestBody ResetPasswordBody
	var err error
	defer utils.HandleError(&err, ctx)

	if err := ctx.BindJSON(&requestBody); err != nil {
		return
	}

	token, err := c.UserService.CreateResetPasswordToken(requestBody.Email)
	if err != nil {
		return
	}

	emailContent := fmt.Sprintf("<p>Otrzymaliśmy prośbę o zresetowanie hasła. Potwierdź prośbę, aby wybrać nowe hasło. Jeśli nie chcesz zmieniać hasła, zignoruj tę wiadomość.</p><p>Link do nadania nowego hasła: %s/nadawanie-hasla#%s</p>", c.Cfg.Server.ClientBaseUrl, token)

	if err = c.Mailer.Send(mailer.MailerMessage{
		Sender:   c.Mailer.GetNoReplySender(),
		Receiver: requestBody.Email,
		Subject:  "Zresetowanie hasła",
		Content:  emailContent,
	}); err != nil {
		return
	}

	ctx.JSON(http.StatusOK, true)
}

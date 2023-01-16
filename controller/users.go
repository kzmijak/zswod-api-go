package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/modules/mailer"
	"github.com/kzmijak/zswod_api_go/services/user"
	"golang.org/x/exp/slices"
)

const (
	ErrPasswordsNotMatching = "ErrPasswordsNotMatching: Passwords should match"
	ErrInvalidRole = "ErrInvalidRole: User has invalid role"
	ErrForbiddenRole = "ErrForbiddenRole: Only website administrators can create users with this role"
	ErrCouldNotParseUuid = "ErrCouldNotParseUuid: Could not parse string to uuid"
)

func (c *Controller) GetAllUsers(ctx *gin.Context) {
	c.log.Trace("Getting users")
	users, err := c.userService.GetAllUsers()

	if err != nil {
		c.log.Error(err)
	}

	ctx.IndentedJSON(http.StatusOK, users)
}

type CreateUserRequest struct {
	user.CreateUserRequest
	PasswordConfirm string `json:"passwordConfirm"`
}

func (c Controller) CreateUser(ctx *gin.Context) {
	var requestBody CreateUserRequest 

	c.log.Trace("Creating user")
	
	if err := ctx.BindJSON(&requestBody); err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	elevatedRoles := []role.Role{
		role.Admin, 
		role.Teacher,
	}

	requestRole, _ := role.FromString(requestBody.Role)
	if slices.Contains(elevatedRoles, requestRole) {
		tokenRole, err := c.ExtractTokenRole(ctx)

		if err != nil || *tokenRole != role.Admin  {
			ctx.JSON(http.StatusForbidden, ErrForbiddenRole)
			return 
		}
	}
	
	if requestBody.Password != requestBody.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, ErrPasswordsNotMatching)
		return
	}

	_, err := c.userService.CreateUser(requestBody.CreateUserRequest)

	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
		return 
	}

	c.log.Trace("Creating user success")

	ctx.IndentedJSON(http.StatusAccepted, nil)
}

func (c *Controller) SignIn(ctx *gin.Context) {
	var requestBody user.SignInRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return;
	}

	response, err := c.userService.SignIn(requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return;
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

func (c *Controller) GetCurrentUser(ctx *gin.Context) {
	token, err := c.ExtractTokenID(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

type ResetPasswordBody struct {
	Email string `json:"email"`
}
func (c *Controller) ResetPassword(ctx *gin.Context) {
	var requestBody ResetPasswordBody
	var err error
	defer utils.HandleError(&err, ctx)

	if err := ctx.BindJSON(&requestBody); err != nil {
		return;
	}

	token, err := c.userService.CreateResetPasswordToken(requestBody.Email)
	if err != nil {
		return;
	}

	emailContent := fmt.Sprintf("<p>Otrzymaliśmy prośbę o zresetowanie hasła. Potwierdź prośbę, aby wybrać nowe hasło. Jeśli nie chcesz zmieniać hasła, zignoruj tę wiadomość.</p><p>Link do nadania nowego hasła: %s/nadawanie-hasla#%s</p>", c.cfg.Server.ClientBaseUrl, token)

	if err = c.mailer.Send(mailer.MailerMessage{
		Sender: c.mailer.GetNoReplySender(),
		Receiver: requestBody.Email,
		Subject: "Zresetowanie hasła",
		Content: emailContent,
	}); err != nil {
		return
	}


	ctx.JSON(http.StatusOK, true)
}


func (c *Controller) VerifyResetPasswordToken(ctx *gin.Context) {
	tokenString := ctx.Query("token")
	var err error
	defer utils.HandleError(&err, ctx)
	
	tokenUuid, err := uuid.Parse(tokenString) 
	if err != nil {
		err = errors.Error(ErrCouldNotParseUuid)
		return
	}

	_, err = c.userService.GetResetPasswordTokenOwner(tokenUuid)
	if err != nil {
		return;
	}

	ctx.JSON(http.StatusOK, true)
}

type SetNewPasswordBody struct {
	Token string `json:"token"`
	Password string `json:"password"`
}
func (c *Controller) SetNewPassword(ctx *gin.Context) {
	var requestBody SetNewPasswordBody
	var err error
	defer utils.HandleError(&err, ctx)

	if err := ctx.BindJSON(&requestBody); err != nil {
		return;
	}

	tokenUuid, err := uuid.Parse(requestBody.Token) 
	if err != nil {
		err = errors.Error(ErrCouldNotParseUuid)
		return
	}

	owner, err := c.userService.GetResetPasswordTokenOwner(tokenUuid)
	if err != nil {
		return
	}

	err = c.userService.SetNewPassword(user.SetNewPasswordParams{
		UserId: owner.ID,
		NewPassword: requestBody.Password,
	})
	if err != nil {
		return
	}

	err = c.userService.InvalidateResetPasswordToken(tokenUuid)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, true)
}
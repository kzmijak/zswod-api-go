package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/services/user"
	"golang.org/x/exp/slices"
)

const (
	ErrPasswordsNotMatching = "ErrPasswordsNotMatching: Passwords should match"
	ErrForbiddenRole = "ErrForbiddenRole: Only website administrators can create users with this role"
)

type CreateUserRequest struct {
	user.CreateUserPayload
	PasswordConfirm string `json:"passwordConfirm"`
}

func (c UserController) CreateUser(ctx *gin.Context) {
	var requestBody CreateUserRequest
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	if err = ctx.BindJSON(&requestBody); err != nil {
		return
	}

	token := utils.ExtractToken(ctx)

	c.Log.Trace("Creating user")


	elevatedRoles := []role.Role{
		role.Admin,
		role.Teacher,
	}

	requestRole, _ := role.FromString(requestBody.Role)
	if slices.Contains(elevatedRoles, requestRole) {
		tokenRole, err := c.JwtService.ExtractTokenRole(token)

		if err != nil || *tokenRole != role.Admin {
			ctx.JSON(http.StatusForbidden, ErrForbiddenRole)
			return
		}
	}

	if requestBody.Password != requestBody.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, ErrPasswordsNotMatching)
		return
	}

	_, err = c.UserService.CreateUser(requestBody.CreateUserPayload, tx)

	if err != nil {
		c.Log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	c.Log.Trace("Creating user success")

	tx.Commit()
	ctx.IndentedJSON(http.StatusAccepted, nil)
}

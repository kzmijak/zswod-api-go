package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
)

type GetAuthenticatedUserResponse struct {
	Id   string `json:"id"`
	Role string `json:"role"`
}

func (c UserController) GetAuthenticatedUser(ctx *gin.Context) {
	var err error
	defer utils.HandleError(&err, ctx)
	token := utils.ExtractToken(ctx)

	roleId, err := c.JwtService.ExtractTokenID(token)
	if err != nil {
		return
	}

	id, err := c.JwtService.ExtractTokenID(token)
	if err != nil {
		return
	}

	response := GetAuthenticatedUserResponse{Id: id.String(), Role: roleId.String()}

	ctx.JSON(http.StatusOK, response)
}
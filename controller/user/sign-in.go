package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/services/user"
)

func (c UserController) SignIn(ctx *gin.Context) {
	var requestBody user.SignInRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := c.UserService.SignIn(requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

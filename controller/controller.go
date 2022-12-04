package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/models"
)


type Controller struct {
	log models.ILogger
	ctx *context.Context
} 

func NewController() *Controller {
	return &Controller{}
}

func (c Controller) WithLogger(log models.ILogger) *Controller  {
	c.log = log
	return &c;
}

func (c Controller) WithContext(ctx *context.Context) *Controller  {
	c.ctx = ctx
	return &c;
}

func (c Controller) Run(host string) {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", c.GetAllUsers)
		}
	}

	router.Run(":" + host)
}

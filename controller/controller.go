package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/models"
	"github.com/kzmijak/zswod_api_go/modules/config"
	"github.com/kzmijak/zswod_api_go/services/jwt"
	"github.com/kzmijak/zswod_api_go/services/user"
)


type Controller struct {
	log models.ILogger
	ctx context.Context
	cfg config.Config

	jwtService jwt.JwtService
	userService user.UserService
} 

func New() Controller {
	return Controller{}
}

func (c Controller) WithLogger(log models.ILogger) Controller  {
	c.log = log
	return c;
}

func (c Controller) WithContext(ctx context.Context) Controller  {
	c.ctx = ctx
	return c;
}

func (c Controller) WithConfig(cfg config.Config) Controller {
	c.cfg = cfg
	return c
}

func (c Controller) Run() {
	router := gin.Default()

	c.jwtService = jwt.New().WithConfig(c.cfg.Auth)
	c.userService = user.New().WithJwtService(c.jwtService).WithContext(c.ctx)

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", c.GetAllUsers)
			users.POST("", c.CreateUser)
			users.POST("sign-in", c.SignIn)
		}
	}

	router.Run(c.cfg.Server.Domain + ":" + c.cfg.Server.Host)
}

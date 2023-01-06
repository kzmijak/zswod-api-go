package controller

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/modules/config"
	"github.com/kzmijak/zswod_api_go/modules/logger"
	"github.com/kzmijak/zswod_api_go/services/article"
	"github.com/kzmijak/zswod_api_go/services/blob"
	"github.com/kzmijak/zswod_api_go/services/image"
	"github.com/kzmijak/zswod_api_go/services/jwt"
	"github.com/kzmijak/zswod_api_go/services/user"
)


type Controller struct {
	log *logger.Logger
	ctx context.Context
	cfg *config.Config

	jwtService *jwt.JwtService
	userService *user.UserService
	blobService blob.BlobService
	articleService article.ArticleService
	imageService image.ImageService
} 

func New() *Controller {
	return &Controller{}
}

func (c *Controller) WithLogger(log *logger.Logger) *Controller  {
	c.log = log
	return c;
}

func (c *Controller) WithContext(ctx context.Context) *Controller  {
	c.ctx = ctx
	return c;
}

func (c *Controller) WithConfig(cfg *config.Config) *Controller {
	c.cfg = cfg
	return c
}

func (c *Controller) Run() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://sporlowd.pl:3000"},
        AllowMethods:     []string{"*"},
        AllowHeaders: []string{"*"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	c.jwtService = jwt.New().WithConfig(c.cfg.Auth)
	c.userService = user.New().WithJwtService(c.jwtService).WithContext(c.ctx)
	c.blobService = blob.New().WithContext(c.ctx)
	c.articleService = article.New().WithContext(c.ctx)
	c.imageService = image.New().WithContext(c.ctx)

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users").Use(c.RequireAuthenticated)
		{
			users.GET("", c.GetAllUsers)
			users.GET("/current", c.GetCurrentUser)
		}
		auth := v1.Group("/auth") 
		{
			auth.POST("sign-up", c.CreateUser)
			auth.POST("sign-in", c.SignIn)	
		}

		blob := v1.Group("/blob")
		{
			blob.GET("/:uuid", c.GetBlobByUuid)
			blob.POST("", c.UploadBlob).Use(c.RequireTeacher)
		}

		article := v1.Group("/article")
		{
			article.POST("/create", c.CreateArticle).Use(c.RequireTeacher)
			article.GET("/:title", c.GetArticleByTitle)
		}
	}

	router.Run(c.cfg.Server.Domain + ":" + c.cfg.Server.Host)
}

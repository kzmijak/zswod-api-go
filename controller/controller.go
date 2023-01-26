package controller

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/modules/config"
	"github.com/kzmijak/zswod_api_go/modules/logger"
	"github.com/kzmijak/zswod_api_go/modules/mailer"
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
	mailer *mailer.Mailer

	jwtService *jwt.JwtService
	userService *user.UserService
	blobService *blob.BlobService
	articleService *article.ArticleService
	imageService *image.ImageService
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

func (c *Controller) WithMailer(mailer *mailer.Mailer) *Controller {
	c.mailer = mailer
	return c
}

func (c *Controller) Run() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"https://sporlowd.pl:3000"},
        AllowMethods:     []string{"*"},
        AllowHeaders: 		[]string{"*"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	c.jwtService = jwt.New().WithConfig(c.cfg.Auth)
	c.userService = user.New().WithJwtService(c.jwtService).WithContext(c.ctx)
	c.blobService = blob.New().WithContext(c.ctx)
	c.imageService = image.New().WithContext(c.ctx).WithBlobService(c.blobService)
	c.articleService = article.New().WithContext(c.ctx).WithImageService(c.imageService)

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users").Use(c.RequireAuthenticated)
		{
			users.GET("", c.GetAllUsers)
			users.GET("/current", c.GetCurrentUser)
		}
		auth := v1.Group("/auth") 
		{
			auth.POST("/sign-up", c.CreateUser)
			auth.POST("/sign-in", c.SignIn)	
			auth.POST("/reset-password", c.ResetPassword)
			auth.GET("/check-reset-password-token", c.VerifyResetPasswordToken)
			auth.POST("/set-new-password", c.SetNewPassword)
		}

		blob := v1.Group("/blob")
		{
			blob.GET("/:uuid", c.GetBlobByUuid)
			blob.GET("", c.GetBlobsList)
			blob.POST("", c.UploadBlob).Use(c.RequireTeacher)
			// blob.GET("", c.GetBlobsList).Use(c.RequireTeacher)
		}

		article := v1.Group("/article")
		{
			article.GET("/:title", c.GetArticleByTitle)
			article.GET("", c.GetArticleHeadersList)
			article.POST("/create", c.CreateArticle).Use(c.RequireTeacher)
			article.PATCH("/update", c.UpdateArticleByTitle).Use(c.RequireTeacher)
		}
	}

	router.RunTLS(c.cfg.Server.Domain + ":" + c.cfg.Server.Host, "cert/sporlowd.pl.crt","cert/sporlowd.pl.key")
}

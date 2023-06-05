package controller

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	articleController "github.com/kzmijak/zswod_api_go/controller/article"
	authController "github.com/kzmijak/zswod_api_go/controller/auth"
	blobController "github.com/kzmijak/zswod_api_go/controller/blob"
	"github.com/kzmijak/zswod_api_go/controller/customPageController"
	galleryController "github.com/kzmijak/zswod_api_go/controller/gallery"
	jwtController "github.com/kzmijak/zswod_api_go/controller/jwt"
	model "github.com/kzmijak/zswod_api_go/controller/model"
	userController "github.com/kzmijak/zswod_api_go/controller/user"
	"github.com/kzmijak/zswod_api_go/modules/config"
	"github.com/kzmijak/zswod_api_go/modules/logger"
	"github.com/kzmijak/zswod_api_go/modules/mailer"
	"github.com/kzmijak/zswod_api_go/services/articleService"
	"github.com/kzmijak/zswod_api_go/services/blobService"
	"github.com/kzmijak/zswod_api_go/services/customPageService"
	"github.com/kzmijak/zswod_api_go/services/galleryService"
	"github.com/kzmijak/zswod_api_go/services/imageService"
	"github.com/kzmijak/zswod_api_go/services/jwt"
	"github.com/kzmijak/zswod_api_go/services/resetPasswordTokenService"
	"github.com/kzmijak/zswod_api_go/services/userService"
)

type Controller struct {
	model.Controller
} 

func New() Controller {
	return Controller{}
}

func (c Controller) WithLogger(log logger.Logger) Controller  {
	c.Log = log
	return c;
}

func (c Controller) WithContext(ctx context.Context) Controller  {
	c.Ctx = ctx
	return c;
}

func (c Controller) WithConfig(cfg config.Config) Controller {
	c.Cfg = cfg
	return c
}

func (c Controller) WithMailer(mailer mailer.Mailer) Controller {
	c.Mailer = mailer
	return c
}

func (c Controller) Run() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
				AllowOrigins:     []string{c.Cfg.Server.ClientBaseUrl},
        AllowMethods:     []string{"*"},
        AllowHeaders: 		[]string{"*"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	c.JwtService = jwt.New().WithConfig(c.Cfg.Auth)
	c.UserService = userService.New(c.Ctx)
	c.BlobService = blobService.New(c.Ctx)
	c.ImageService = imageService.New().WithContext(c.Ctx)
	c.ArticleService = articleService.New(c.Ctx)
	c.GalleryService = galleryService.New(c.Ctx)
	c.CustomPageService = customPageService.New(c.Ctx)
	c.ResetPasswordTokenService = resetPasswordTokenService.New(c.Ctx)
	
	jwtController.New(c.Controller)
	jc := jwtController.New(c.Controller)

	ac := authController.New(c.Controller)

	uc := userController.New(c.Controller)
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users").Use(jc.RequireAuthenticated)
		{
			users.GET("", uc.GetAllUsers)
			users.GET("/current", uc.GetAuthenticatedUser)
		}
		auth := v1.Group("/auth") 
		{
			auth.POST("/sign-up", uc.CreateUser)
			auth.POST("/sign-in", ac.SignIn)	
			auth.POST("/reset-password", uc.ResetPassword)
			auth.GET("/check-reset-password-token", uc.VerifyResetPasswordToken)
			auth.POST("/set-new-password", uc.SetNewPassword)
		}

		bc := blobController.New(c.Controller)
		blob := v1.Group("/blob")
		{
			blob.GET("/:uuid", bc.GetBlobByUuid)
			blob.Use(jc.RequireTeacher)
			blob.POST("", bc.UploadBlob)
			blob.GET("", bc.GetBlobsList)
		}

		ac := articleController.New(c.Controller)
		article := v1.Group("/article")
		{
			article.GET("/:title", ac.GetArticleByTitle)
			article.GET("", ac.GetArticleHeadersList)
			article.POST("/create", ac.CreateArticle).Use(jc.RequireTeacher)
			article.PUT("/update", ac.UpdateArticle).Use(jc.RequireTeacher)
			article.DELETE("/:articleId", ac.RemoveArticle).Use(jc.RequireTeacher)
		}

		gc := galleryController.New(c.Controller)
		gallery := v1.Group("/gallery")
		{
			gallery.GET("", gc.GetGalleryHeadersList)
			gallery.GET("/:galleryId", gc.GetGalleryById)
		}

		cpc := customPageController.New(c.Controller)
		customPage := v1.Group("/customPages")
		{
			customPage.GET("", cpc.GetCustomPageHeaders)
			customPage.GET("/:section/:title", cpc.GetCustomPageBySectionAndTitle)
			customPage.Use(jc.RequireTeacher)
			customPage.POST("", cpc.CreateCustomPage)
			customPage.PUT("", cpc.UpdateCustomPage)
			customPage.DELETE("/:id", cpc.DeleteCustomPage)
		}
	}

	router.RunTLS(c.Cfg.Server.Domain + ":" + c.Cfg.Server.Host, "cert/sporlowd.pl.crt","cert/sporlowd.pl.key")
}

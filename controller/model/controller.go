package controller

import (
	"context"

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
	Log    logger.Logger
	Ctx    context.Context
	Cfg    config.Config
	Mailer mailer.Mailer

	JwtService     jwt.JwtService
	UserService    userService.UserService
	BlobService    blobService.BlobService
	ArticleService articleService.ArticleService
	ImageService   imageService.ImageService
	GalleryService galleryService.GalleryService
	CustomPageService customPageService.CustomPageService
	ResetPasswordTokenService resetPasswordTokenService.ResetPasswordTokenService
}
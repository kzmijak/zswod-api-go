package controller

import (
	"context"

	"github.com/kzmijak/zswod_api_go/modules/config"
	"github.com/kzmijak/zswod_api_go/modules/logger"
	"github.com/kzmijak/zswod_api_go/modules/mailer"
	"github.com/kzmijak/zswod_api_go/services/articleService"
	"github.com/kzmijak/zswod_api_go/services/blob"
	"github.com/kzmijak/zswod_api_go/services/galleryService"
	"github.com/kzmijak/zswod_api_go/services/image"
	"github.com/kzmijak/zswod_api_go/services/jwt"
	"github.com/kzmijak/zswod_api_go/services/user"
)

type Controller struct {
	Log    logger.Logger
	Ctx    context.Context
	Cfg    config.Config
	Mailer mailer.Mailer

	JwtService     jwt.JwtService
	UserService    user.UserService
	BlobService    blob.BlobService
	ArticleService articleService.ArticleService
	ImageService   image.ImageService
	GalleryService galleryService.GalleryService
}
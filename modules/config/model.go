package config

import (
	"github.com/kzmijak/zswod_api_go/modules/controller"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/logger"
)

type Config struct {
	Database database.DatabaseConfig `json:"database"`
	Logger logger.LoggerConfig `json:"logger"`
	Server controller.ControllerConfig `json:"server"`
}
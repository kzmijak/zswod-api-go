package logger

import (
	"context"
	"os"

	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/sirupsen/logrus"
)

const (
	ErrInvalidFile = "err_invalid_file: Failed to set logger output file"
)

type Logger struct {
	*logrus.Logger
}

func Initialize(ctx context.Context, cfg LoggerConfig) (*Logger, error) {
	var lgr = &Logger{
		logrus.New(),
	}
	

	lgr.SetLevel(logrus.TraceLevel)

	file, err := os.OpenFile(cfg.FilePath, os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		return nil, errors.Error(ErrInvalidFile)
	}
	lgr.SetOutput(file)
	lgr.SetFormatter(&logrus.JSONFormatter{})


	return lgr, nil
}

func (lgr Logger) GetLevel() string {
	return lgr.Level.String()
}
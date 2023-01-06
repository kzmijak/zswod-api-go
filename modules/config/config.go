package config

import (
	"encoding/json"
	"os"

	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const baseConfigPath = "config/config.json"

const (
	ErrFileNotFound = errors.Error("ErrFileNotFound: Provided filename does not lead to any file")
	ErrCannotParseJson = errors.Error("ErrCannotParseJson: Provided object failed to be parsed from JSON")
)

func Initialize() (*Config, error) {
	cfg := &Config{}

	data, err := os.ReadFile(baseConfigPath)
	if err != nil {
		return nil, ErrFileNotFound
	}

	if err := json.Unmarshal([]byte(data), cfg); err != nil {
		return nil, ErrCannotParseJson
	}

	return cfg, nil
}
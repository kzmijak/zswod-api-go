package config

import (
	"encoding/json"
	"os"

	"github.com/kzmijak/zswod_api_go/internal/errors"
)

const baseConfigPath = "config/config.json"

const (
	ErrFileNotFound = errors.Error("err_file_not_found: Provided filename does not lead to any file")
	ErrCannotParseJson = errors.Error("err_cannot_parse_json: Provided object failed to be parsed from JSON")
)

func LoadJson(cfg *Config) error {
	data, err := os.ReadFile(baseConfigPath)
	if err != nil {
		return ErrFileNotFound
	}

	if err := json.Unmarshal([]byte(data), cfg); err != nil {
		return ErrCannotParseJson
	}

	return nil
}
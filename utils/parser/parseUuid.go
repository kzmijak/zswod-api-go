package parser

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrUuidParseFailed = "Could not parse input string to UUID"
)

func ParseUuid(uuidStr string) (uuid.UUID, error) {
	uuidParsed, err := uuid.Parse(uuidStr)
	if err != nil {
		return uuid.Nil, errors.Error(ErrUuidParseFailed)
	}

	return uuidParsed, nil
}
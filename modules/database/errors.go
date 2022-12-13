package database

import "github.com/kzmijak/zswod_api_go/modules/errors"

const (
	ErrConnectionFailed   = errors.Error("err_connection_failed: Failed opening connection to MySQL")
	ErrSchemaCreationFail = errors.Error("err_schema_creation_fail: Failed creating schema resource")
)
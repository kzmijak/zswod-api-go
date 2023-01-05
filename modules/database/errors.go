package database

import "github.com/kzmijak/zswod_api_go/modules/errors"

const (
	ErrConnectionFailed   = errors.Error("ErrConnectionFailed: Failed opening connection to MySQL")
	ErrSchemaCreationFail = errors.Error("ErrSchemaCreationFail: Failed creating schema resource")
	ErrCouldNotQuery = errors.Error("ErrCouldNotQuery: Failed query db database")
	ErrRolesSeedFail = errors.Error("ErrRolesSeedFail: Failed to seed roles")
)
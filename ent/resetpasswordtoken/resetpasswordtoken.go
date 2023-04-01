// Code generated by ent, DO NOT EDIT.

package resetpasswordtoken

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the resetpasswordtoken type in the database.
	Label = "reset_password_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldExpiryTime holds the string denoting the expirytime field in the database.
	FieldExpiryTime = "expiry_time"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the resetpasswordtoken in the database.
	Table = "reset_password_tokens"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "reset_password_tokens"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_reset_password_token"
)

// Columns holds all SQL columns for resetpasswordtoken fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldExpiryTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "reset_password_tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_reset_password_token",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

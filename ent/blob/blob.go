// Code generated by ent, DO NOT EDIT.

package blob

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the blob type in the database.
	Label = "blob"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldBlob holds the string denoting the blob field in the database.
	FieldBlob = "blob"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContentType holds the string denoting the contenttype field in the database.
	FieldContentType = "content_type"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIsPublic holds the string denoting the ispublic field in the database.
	FieldIsPublic = "is_public"
	// Table holds the table name of the blob in the database.
	Table = "blobs"
)

// Columns holds all SQL columns for blob fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldBlob,
	FieldTitle,
	FieldContentType,
	FieldType,
	FieldIsPublic,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultIsPublic holds the default value on creation for the "isPublic" field.
	DefaultIsPublic bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// TypePicture is the default value of the Type enum.
const DefaultType = TypePicture

// Type values.
const (
	TypePicture    Type = "Picture"
	TypeAttachment Type = "Attachment"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypePicture, TypeAttachment:
		return nil
	default:
		return fmt.Errorf("blob: invalid enum value for type field: %q", _type)
	}
}

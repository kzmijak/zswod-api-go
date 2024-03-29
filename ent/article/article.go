// Code generated by ent, DO NOT EDIT.

package article

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the article type in the database.
	Label = "article"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldTitleNormalized holds the string denoting the titlenormalized field in the database.
	FieldTitleNormalized = "title_normalized"
	// FieldShort holds the string denoting the short field in the database.
	FieldShort = "short"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeGallery holds the string denoting the gallery edge name in mutations.
	EdgeGallery = "gallery"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeAttachments holds the string denoting the attachments edge name in mutations.
	EdgeAttachments = "attachments"
	// Table holds the table name of the article in the database.
	Table = "articles"
	// GalleryTable is the table that holds the gallery relation/edge.
	GalleryTable = "galleries"
	// GalleryInverseTable is the table name for the Gallery entity.
	// It exists in this package in order to avoid circular dependency with the "gallery" package.
	GalleryInverseTable = "galleries"
	// GalleryColumn is the table column denoting the gallery relation/edge.
	GalleryColumn = "article_gallery"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "articles"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "user_articles"
	// AttachmentsTable is the table that holds the attachments relation/edge.
	AttachmentsTable = "attachments"
	// AttachmentsInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentsInverseTable = "attachments"
	// AttachmentsColumn is the table column denoting the attachments relation/edge.
	AttachmentsColumn = "article_attachments"
)

// Columns holds all SQL columns for article fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldTitle,
	FieldTitleNormalized,
	FieldShort,
	FieldContent,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "articles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_articles",
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// ShortValidator is a validator for the "short" field. It is called by the builders before save.
	ShortValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusDraft     Status = "Draft"
	StatusReview    Status = "Review"
	StatusPublished Status = "Published"
	StatusRemoved   Status = "Removed"
	StatusUnknown   Status = "Unknown"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusDraft, StatusReview, StatusPublished, StatusRemoved, StatusUnknown:
		return nil
	default:
		return fmt.Errorf("article: invalid enum value for status field: %q", s)
	}
}

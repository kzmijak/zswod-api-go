// Code generated by ent, DO NOT EDIT.

package image

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the image type in the database.
	Label = "image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldAlt holds the string denoting the alt field in the database.
	FieldAlt = "alt"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldBlobId holds the string denoting the blobid field in the database.
	FieldBlobId = "blob_id"
	// EdgeGallery holds the string denoting the gallery edge name in mutations.
	EdgeGallery = "gallery"
	// EdgeBlob holds the string denoting the blob edge name in mutations.
	EdgeBlob = "blob"
	// Table holds the table name of the image in the database.
	Table = "images"
	// GalleryTable is the table that holds the gallery relation/edge.
	GalleryTable = "images"
	// GalleryInverseTable is the table name for the Gallery entity.
	// It exists in this package in order to avoid circular dependency with the "gallery" package.
	GalleryInverseTable = "galleries"
	// GalleryColumn is the table column denoting the gallery relation/edge.
	GalleryColumn = "gallery_images"
	// BlobTable is the table that holds the blob relation/edge.
	BlobTable = "images"
	// BlobInverseTable is the table name for the Blob entity.
	// It exists in this package in order to avoid circular dependency with the "blob" package.
	BlobInverseTable = "blobs"
	// BlobColumn is the table column denoting the blob relation/edge.
	BlobColumn = "blob_id"
)

// Columns holds all SQL columns for image fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldAlt,
	FieldOrder,
	FieldBlobId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"gallery_images",
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
	// OrderValidator is a validator for the "Order" field. It is called by the builders before save.
	OrderValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

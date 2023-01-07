// Code generated by ent, DO NOT EDIT.

package blob

const (
	// Label holds the string label denoting the blob type in the database.
	Label = "blob"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBlob holds the string denoting the blob field in the database.
	FieldBlob = "blob"
	// FieldContentType holds the string denoting the contenttype field in the database.
	FieldContentType = "content_type"
	// Table holds the table name of the blob in the database.
	Table = "blobs"
)

// Columns holds all SQL columns for blob fields.
var Columns = []string{
	FieldID,
	FieldBlob,
	FieldContentType,
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

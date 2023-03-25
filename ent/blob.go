// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/blob"
)

// Blob is the model entity for the Blob schema.
type Blob struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Blob holds the value of the "blob" field.
	Blob []byte `json:"blob,omitempty"`
	// ContentType holds the value of the "contentType" field.
	ContentType string `json:"contentType,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blob) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blob.FieldBlob:
			values[i] = new([]byte)
		case blob.FieldContentType:
			values[i] = new(sql.NullString)
		case blob.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case blob.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Blob", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blob fields.
func (b *Blob) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blob.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case blob.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				b.CreateTime = value.Time
			}
		case blob.FieldBlob:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field blob", values[i])
			} else if value != nil {
				b.Blob = *value
			}
		case blob.FieldContentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contentType", values[i])
			} else if value.Valid {
				b.ContentType = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Blob.
// Note that you need to call Blob.Unwrap() before calling this method if this Blob
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Blob) Update() *BlobUpdateOne {
	return (&BlobClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Blob entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Blob) Unwrap() *Blob {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blob is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Blob) String() string {
	var builder strings.Builder
	builder.WriteString("Blob(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("create_time=")
	builder.WriteString(b.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("blob=")
	builder.WriteString(fmt.Sprintf("%v", b.Blob))
	builder.WriteString(", ")
	builder.WriteString("contentType=")
	builder.WriteString(b.ContentType)
	builder.WriteByte(')')
	return builder.String()
}

// Blobs is a parsable slice of Blob.
type Blobs []*Blob

func (b Blobs) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}

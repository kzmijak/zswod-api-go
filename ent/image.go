// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/ent/image"
)

// Image is the model entity for the Image schema.
type Image struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Alt holds the value of the "alt" field.
	Alt string `json:"alt,omitempty"`
	// IsPreview holds the value of the "isPreview" field.
	IsPreview bool `json:"isPreview,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImageQuery when eager-loading is set.
	Edges               ImageEdges `json:"edges"`
	article_images      *uuid.UUID
	blob_article_images *uuid.UUID
}

// ImageEdges holds the relations/edges for other nodes in the graph.
type ImageEdges struct {
	// Article holds the value of the article edge.
	Article *Article `json:"article,omitempty"`
	// Blob holds the value of the blob edge.
	Blob *Blob `json:"blob,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ArticleOrErr returns the Article value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageEdges) ArticleOrErr() (*Article, error) {
	if e.loadedTypes[0] {
		if e.Article == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: article.Label}
		}
		return e.Article, nil
	}
	return nil, &NotLoadedError{edge: "article"}
}

// BlobOrErr returns the Blob value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageEdges) BlobOrErr() (*Blob, error) {
	if e.loadedTypes[1] {
		if e.Blob == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: blob.Label}
		}
		return e.Blob, nil
	}
	return nil, &NotLoadedError{edge: "blob"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Image) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case image.FieldIsPreview:
			values[i] = new(sql.NullBool)
		case image.FieldTitle, image.FieldAlt:
			values[i] = new(sql.NullString)
		case image.FieldID:
			values[i] = new(uuid.UUID)
		case image.ForeignKeys[0]: // article_images
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case image.ForeignKeys[1]: // blob_article_images
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Image", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Image fields.
func (i *Image) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case image.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case image.FieldTitle:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[j])
			} else if value.Valid {
				i.Title = value.String
			}
		case image.FieldAlt:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alt", values[j])
			} else if value.Valid {
				i.Alt = value.String
			}
		case image.FieldIsPreview:
			if value, ok := values[j].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isPreview", values[j])
			} else if value.Valid {
				i.IsPreview = value.Bool
			}
		case image.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field article_images", values[j])
			} else if value.Valid {
				i.article_images = new(uuid.UUID)
				*i.article_images = *value.S.(*uuid.UUID)
			}
		case image.ForeignKeys[1]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field blob_article_images", values[j])
			} else if value.Valid {
				i.blob_article_images = new(uuid.UUID)
				*i.blob_article_images = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryArticle queries the "article" edge of the Image entity.
func (i *Image) QueryArticle() *ArticleQuery {
	return (&ImageClient{config: i.config}).QueryArticle(i)
}

// QueryBlob queries the "blob" edge of the Image entity.
func (i *Image) QueryBlob() *BlobQuery {
	return (&ImageClient{config: i.config}).QueryBlob(i)
}

// Update returns a builder for updating this Image.
// Note that you need to call Image.Unwrap() before calling this method if this Image
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Image) Update() *ImageUpdateOne {
	return (&ImageClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Image entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Image) Unwrap() *Image {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Image is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Image) String() string {
	var builder strings.Builder
	builder.WriteString("Image(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("title=")
	builder.WriteString(i.Title)
	builder.WriteString(", ")
	builder.WriteString("alt=")
	builder.WriteString(i.Alt)
	builder.WriteString(", ")
	builder.WriteString("isPreview=")
	builder.WriteString(fmt.Sprintf("%v", i.IsPreview))
	builder.WriteByte(')')
	return builder.String()
}

// Images is a parsable slice of Image.
type Images []*Image

func (i Images) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/ent/articletitleguid"
)

// ArticleTitleGuid is the model entity for the ArticleTitleGuid schema.
type ArticleTitleGuid struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// TitleNormalized holds the value of the "titleNormalized" field.
	TitleNormalized string `json:"titleNormalized,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArticleTitleGuidQuery when eager-loading is set.
	Edges                    ArticleTitleGuidEdges `json:"edges"`
	article_title_normalized *uuid.UUID
}

// ArticleTitleGuidEdges holds the relations/edges for other nodes in the graph.
type ArticleTitleGuidEdges struct {
	// Article holds the value of the article edge.
	Article *Article `json:"article,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ArticleOrErr returns the Article value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArticleTitleGuidEdges) ArticleOrErr() (*Article, error) {
	if e.loadedTypes[0] {
		if e.Article == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: article.Label}
		}
		return e.Article, nil
	}
	return nil, &NotLoadedError{edge: "article"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ArticleTitleGuid) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case articletitleguid.FieldTitleNormalized:
			values[i] = new(sql.NullString)
		case articletitleguid.FieldID:
			values[i] = new(uuid.UUID)
		case articletitleguid.ForeignKeys[0]: // article_title_normalized
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ArticleTitleGuid", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ArticleTitleGuid fields.
func (atg *ArticleTitleGuid) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case articletitleguid.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				atg.ID = *value
			}
		case articletitleguid.FieldTitleNormalized:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field titleNormalized", values[i])
			} else if value.Valid {
				atg.TitleNormalized = value.String
			}
		case articletitleguid.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field article_title_normalized", values[i])
			} else if value.Valid {
				atg.article_title_normalized = new(uuid.UUID)
				*atg.article_title_normalized = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryArticle queries the "article" edge of the ArticleTitleGuid entity.
func (atg *ArticleTitleGuid) QueryArticle() *ArticleQuery {
	return (&ArticleTitleGuidClient{config: atg.config}).QueryArticle(atg)
}

// Update returns a builder for updating this ArticleTitleGuid.
// Note that you need to call ArticleTitleGuid.Unwrap() before calling this method if this ArticleTitleGuid
// was returned from a transaction, and the transaction was committed or rolled back.
func (atg *ArticleTitleGuid) Update() *ArticleTitleGuidUpdateOne {
	return (&ArticleTitleGuidClient{config: atg.config}).UpdateOne(atg)
}

// Unwrap unwraps the ArticleTitleGuid entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (atg *ArticleTitleGuid) Unwrap() *ArticleTitleGuid {
	_tx, ok := atg.config.driver.(*txDriver)
	if !ok {
		panic("ent: ArticleTitleGuid is not a transactional entity")
	}
	atg.config.driver = _tx.drv
	return atg
}

// String implements the fmt.Stringer.
func (atg *ArticleTitleGuid) String() string {
	var builder strings.Builder
	builder.WriteString("ArticleTitleGuid(")
	builder.WriteString(fmt.Sprintf("id=%v, ", atg.ID))
	builder.WriteString("titleNormalized=")
	builder.WriteString(atg.TitleNormalized)
	builder.WriteByte(')')
	return builder.String()
}

// ArticleTitleGuids is a parsable slice of ArticleTitleGuid.
type ArticleTitleGuids []*ArticleTitleGuid

func (atg ArticleTitleGuids) config(cfg config) {
	for _i := range atg {
		atg[_i].config = cfg
	}
}

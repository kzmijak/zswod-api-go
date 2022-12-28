// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/ent/articletitleguid"
	"github.com/kzmijak/zswod_api_go/ent/predicate"
)

// ArticleTitleGuidQuery is the builder for querying ArticleTitleGuid entities.
type ArticleTitleGuidQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.ArticleTitleGuid
	withArticle *ArticleQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArticleTitleGuidQuery builder.
func (atgq *ArticleTitleGuidQuery) Where(ps ...predicate.ArticleTitleGuid) *ArticleTitleGuidQuery {
	atgq.predicates = append(atgq.predicates, ps...)
	return atgq
}

// Limit adds a limit step to the query.
func (atgq *ArticleTitleGuidQuery) Limit(limit int) *ArticleTitleGuidQuery {
	atgq.limit = &limit
	return atgq
}

// Offset adds an offset step to the query.
func (atgq *ArticleTitleGuidQuery) Offset(offset int) *ArticleTitleGuidQuery {
	atgq.offset = &offset
	return atgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (atgq *ArticleTitleGuidQuery) Unique(unique bool) *ArticleTitleGuidQuery {
	atgq.unique = &unique
	return atgq
}

// Order adds an order step to the query.
func (atgq *ArticleTitleGuidQuery) Order(o ...OrderFunc) *ArticleTitleGuidQuery {
	atgq.order = append(atgq.order, o...)
	return atgq
}

// QueryArticle chains the current query on the "article" edge.
func (atgq *ArticleTitleGuidQuery) QueryArticle() *ArticleQuery {
	query := &ArticleQuery{config: atgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(articletitleguid.Table, articletitleguid.FieldID, selector),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, articletitleguid.ArticleTable, articletitleguid.ArticleColumn),
		)
		fromU = sqlgraph.SetNeighbors(atgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ArticleTitleGuid entity from the query.
// Returns a *NotFoundError when no ArticleTitleGuid was found.
func (atgq *ArticleTitleGuidQuery) First(ctx context.Context) (*ArticleTitleGuid, error) {
	nodes, err := atgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{articletitleguid.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (atgq *ArticleTitleGuidQuery) FirstX(ctx context.Context) *ArticleTitleGuid {
	node, err := atgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ArticleTitleGuid ID from the query.
// Returns a *NotFoundError when no ArticleTitleGuid ID was found.
func (atgq *ArticleTitleGuidQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{articletitleguid.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (atgq *ArticleTitleGuidQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := atgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ArticleTitleGuid entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ArticleTitleGuid entity is found.
// Returns a *NotFoundError when no ArticleTitleGuid entities are found.
func (atgq *ArticleTitleGuidQuery) Only(ctx context.Context) (*ArticleTitleGuid, error) {
	nodes, err := atgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{articletitleguid.Label}
	default:
		return nil, &NotSingularError{articletitleguid.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (atgq *ArticleTitleGuidQuery) OnlyX(ctx context.Context) *ArticleTitleGuid {
	node, err := atgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ArticleTitleGuid ID in the query.
// Returns a *NotSingularError when more than one ArticleTitleGuid ID is found.
// Returns a *NotFoundError when no entities are found.
func (atgq *ArticleTitleGuidQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{articletitleguid.Label}
	default:
		err = &NotSingularError{articletitleguid.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (atgq *ArticleTitleGuidQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := atgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ArticleTitleGuids.
func (atgq *ArticleTitleGuidQuery) All(ctx context.Context) ([]*ArticleTitleGuid, error) {
	if err := atgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return atgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (atgq *ArticleTitleGuidQuery) AllX(ctx context.Context) []*ArticleTitleGuid {
	nodes, err := atgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ArticleTitleGuid IDs.
func (atgq *ArticleTitleGuidQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := atgq.Select(articletitleguid.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (atgq *ArticleTitleGuidQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := atgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (atgq *ArticleTitleGuidQuery) Count(ctx context.Context) (int, error) {
	if err := atgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return atgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (atgq *ArticleTitleGuidQuery) CountX(ctx context.Context) int {
	count, err := atgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (atgq *ArticleTitleGuidQuery) Exist(ctx context.Context) (bool, error) {
	if err := atgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return atgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (atgq *ArticleTitleGuidQuery) ExistX(ctx context.Context) bool {
	exist, err := atgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArticleTitleGuidQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (atgq *ArticleTitleGuidQuery) Clone() *ArticleTitleGuidQuery {
	if atgq == nil {
		return nil
	}
	return &ArticleTitleGuidQuery{
		config:      atgq.config,
		limit:       atgq.limit,
		offset:      atgq.offset,
		order:       append([]OrderFunc{}, atgq.order...),
		predicates:  append([]predicate.ArticleTitleGuid{}, atgq.predicates...),
		withArticle: atgq.withArticle.Clone(),
		// clone intermediate query.
		sql:    atgq.sql.Clone(),
		path:   atgq.path,
		unique: atgq.unique,
	}
}

// WithArticle tells the query-builder to eager-load the nodes that are connected to
// the "article" edge. The optional arguments are used to configure the query builder of the edge.
func (atgq *ArticleTitleGuidQuery) WithArticle(opts ...func(*ArticleQuery)) *ArticleTitleGuidQuery {
	query := &ArticleQuery{config: atgq.config}
	for _, opt := range opts {
		opt(query)
	}
	atgq.withArticle = query
	return atgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TitleNormalized string `json:"title_normalized,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ArticleTitleGuid.Query().
//		GroupBy(articletitleguid.FieldTitleNormalized).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (atgq *ArticleTitleGuidQuery) GroupBy(field string, fields ...string) *ArticleTitleGuidGroupBy {
	grbuild := &ArticleTitleGuidGroupBy{config: atgq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := atgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return atgq.sqlQuery(ctx), nil
	}
	grbuild.label = articletitleguid.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TitleNormalized string `json:"title_normalized,omitempty"`
//	}
//
//	client.ArticleTitleGuid.Query().
//		Select(articletitleguid.FieldTitleNormalized).
//		Scan(ctx, &v)
func (atgq *ArticleTitleGuidQuery) Select(fields ...string) *ArticleTitleGuidSelect {
	atgq.fields = append(atgq.fields, fields...)
	selbuild := &ArticleTitleGuidSelect{ArticleTitleGuidQuery: atgq}
	selbuild.label = articletitleguid.Label
	selbuild.flds, selbuild.scan = &atgq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a ArticleTitleGuidSelect configured with the given aggregations.
func (atgq *ArticleTitleGuidQuery) Aggregate(fns ...AggregateFunc) *ArticleTitleGuidSelect {
	return atgq.Select().Aggregate(fns...)
}

func (atgq *ArticleTitleGuidQuery) prepareQuery(ctx context.Context) error {
	for _, f := range atgq.fields {
		if !articletitleguid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if atgq.path != nil {
		prev, err := atgq.path(ctx)
		if err != nil {
			return err
		}
		atgq.sql = prev
	}
	return nil
}

func (atgq *ArticleTitleGuidQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ArticleTitleGuid, error) {
	var (
		nodes       = []*ArticleTitleGuid{}
		withFKs     = atgq.withFKs
		_spec       = atgq.querySpec()
		loadedTypes = [1]bool{
			atgq.withArticle != nil,
		}
	)
	if atgq.withArticle != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, articletitleguid.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ArticleTitleGuid).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ArticleTitleGuid{config: atgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, atgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := atgq.withArticle; query != nil {
		if err := atgq.loadArticle(ctx, query, nodes, nil,
			func(n *ArticleTitleGuid, e *Article) { n.Edges.Article = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (atgq *ArticleTitleGuidQuery) loadArticle(ctx context.Context, query *ArticleQuery, nodes []*ArticleTitleGuid, init func(*ArticleTitleGuid), assign func(*ArticleTitleGuid, *Article)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ArticleTitleGuid)
	for i := range nodes {
		if nodes[i].article_title_normalized == nil {
			continue
		}
		fk := *nodes[i].article_title_normalized
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(article.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "article_title_normalized" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (atgq *ArticleTitleGuidQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := atgq.querySpec()
	_spec.Node.Columns = atgq.fields
	if len(atgq.fields) > 0 {
		_spec.Unique = atgq.unique != nil && *atgq.unique
	}
	return sqlgraph.CountNodes(ctx, atgq.driver, _spec)
}

func (atgq *ArticleTitleGuidQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := atgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (atgq *ArticleTitleGuidQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   articletitleguid.Table,
			Columns: articletitleguid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: articletitleguid.FieldID,
			},
		},
		From:   atgq.sql,
		Unique: true,
	}
	if unique := atgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := atgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, articletitleguid.FieldID)
		for i := range fields {
			if fields[i] != articletitleguid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := atgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := atgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := atgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := atgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (atgq *ArticleTitleGuidQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(atgq.driver.Dialect())
	t1 := builder.Table(articletitleguid.Table)
	columns := atgq.fields
	if len(columns) == 0 {
		columns = articletitleguid.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if atgq.sql != nil {
		selector = atgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if atgq.unique != nil && *atgq.unique {
		selector.Distinct()
	}
	for _, p := range atgq.predicates {
		p(selector)
	}
	for _, p := range atgq.order {
		p(selector)
	}
	if offset := atgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := atgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ArticleTitleGuidGroupBy is the group-by builder for ArticleTitleGuid entities.
type ArticleTitleGuidGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atggb *ArticleTitleGuidGroupBy) Aggregate(fns ...AggregateFunc) *ArticleTitleGuidGroupBy {
	atggb.fns = append(atggb.fns, fns...)
	return atggb
}

// Scan applies the group-by query and scans the result into the given value.
func (atggb *ArticleTitleGuidGroupBy) Scan(ctx context.Context, v any) error {
	query, err := atggb.path(ctx)
	if err != nil {
		return err
	}
	atggb.sql = query
	return atggb.sqlScan(ctx, v)
}

func (atggb *ArticleTitleGuidGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range atggb.fields {
		if !articletitleguid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := atggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (atggb *ArticleTitleGuidGroupBy) sqlQuery() *sql.Selector {
	selector := atggb.sql.Select()
	aggregation := make([]string, 0, len(atggb.fns))
	for _, fn := range atggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(atggb.fields)+len(atggb.fns))
		for _, f := range atggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(atggb.fields...)...)
}

// ArticleTitleGuidSelect is the builder for selecting fields of ArticleTitleGuid entities.
type ArticleTitleGuidSelect struct {
	*ArticleTitleGuidQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (atgs *ArticleTitleGuidSelect) Aggregate(fns ...AggregateFunc) *ArticleTitleGuidSelect {
	atgs.fns = append(atgs.fns, fns...)
	return atgs
}

// Scan applies the selector query and scans the result into the given value.
func (atgs *ArticleTitleGuidSelect) Scan(ctx context.Context, v any) error {
	if err := atgs.prepareQuery(ctx); err != nil {
		return err
	}
	atgs.sql = atgs.ArticleTitleGuidQuery.sqlQuery(ctx)
	return atgs.sqlScan(ctx, v)
}

func (atgs *ArticleTitleGuidSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(atgs.fns))
	for _, fn := range atgs.fns {
		aggregation = append(aggregation, fn(atgs.sql))
	}
	switch n := len(*atgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		atgs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		atgs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := atgs.sql.Query()
	if err := atgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

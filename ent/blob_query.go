// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/attachment"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/ent/image"
	"github.com/kzmijak/zswod_api_go/ent/predicate"
)

// BlobQuery is the builder for querying Blob entities.
type BlobQuery struct {
	config
	limit           *int
	offset          *int
	unique          *bool
	order           []OrderFunc
	fields          []string
	predicates      []predicate.Blob
	withAttachments *AttachmentQuery
	withImages      *ImageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlobQuery builder.
func (bq *BlobQuery) Where(ps ...predicate.Blob) *BlobQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BlobQuery) Limit(limit int) *BlobQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BlobQuery) Offset(offset int) *BlobQuery {
	bq.offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BlobQuery) Unique(unique bool) *BlobQuery {
	bq.unique = &unique
	return bq
}

// Order adds an order step to the query.
func (bq *BlobQuery) Order(o ...OrderFunc) *BlobQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryAttachments chains the current query on the "attachments" edge.
func (bq *BlobQuery) QueryAttachments() *AttachmentQuery {
	query := &AttachmentQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blob.Table, blob.FieldID, selector),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, blob.AttachmentsTable, blob.AttachmentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryImages chains the current query on the "images" edge.
func (bq *BlobQuery) QueryImages() *ImageQuery {
	query := &ImageQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blob.Table, blob.FieldID, selector),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, blob.ImagesTable, blob.ImagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Blob entity from the query.
// Returns a *NotFoundError when no Blob was found.
func (bq *BlobQuery) First(ctx context.Context) (*Blob, error) {
	nodes, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{blob.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BlobQuery) FirstX(ctx context.Context) *Blob {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Blob ID from the query.
// Returns a *NotFoundError when no Blob ID was found.
func (bq *BlobQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blob.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BlobQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Blob entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Blob entity is found.
// Returns a *NotFoundError when no Blob entities are found.
func (bq *BlobQuery) Only(ctx context.Context) (*Blob, error) {
	nodes, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{blob.Label}
	default:
		return nil, &NotSingularError{blob.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BlobQuery) OnlyX(ctx context.Context) *Blob {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Blob ID in the query.
// Returns a *NotSingularError when more than one Blob ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BlobQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blob.Label}
	default:
		err = &NotSingularError{blob.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BlobQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Blobs.
func (bq *BlobQuery) All(ctx context.Context) ([]*Blob, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bq *BlobQuery) AllX(ctx context.Context) []*Blob {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Blob IDs.
func (bq *BlobQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := bq.Select(blob.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BlobQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BlobQuery) Count(ctx context.Context) (int, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BlobQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BlobQuery) Exist(ctx context.Context) (bool, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BlobQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlobQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BlobQuery) Clone() *BlobQuery {
	if bq == nil {
		return nil
	}
	return &BlobQuery{
		config:          bq.config,
		limit:           bq.limit,
		offset:          bq.offset,
		order:           append([]OrderFunc{}, bq.order...),
		predicates:      append([]predicate.Blob{}, bq.predicates...),
		withAttachments: bq.withAttachments.Clone(),
		withImages:      bq.withImages.Clone(),
		// clone intermediate query.
		sql:    bq.sql.Clone(),
		path:   bq.path,
		unique: bq.unique,
	}
}

// WithAttachments tells the query-builder to eager-load the nodes that are connected to
// the "attachments" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BlobQuery) WithAttachments(opts ...func(*AttachmentQuery)) *BlobQuery {
	query := &AttachmentQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withAttachments = query
	return bq
}

// WithImages tells the query-builder to eager-load the nodes that are connected to
// the "images" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BlobQuery) WithImages(opts ...func(*ImageQuery)) *BlobQuery {
	query := &ImageQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withImages = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Blob.Query().
//		GroupBy(blob.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BlobQuery) GroupBy(field string, fields ...string) *BlobGroupBy {
	grbuild := &BlobGroupBy{config: bq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(ctx), nil
	}
	grbuild.label = blob.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Blob.Query().
//		Select(blob.FieldCreateTime).
//		Scan(ctx, &v)
func (bq *BlobQuery) Select(fields ...string) *BlobSelect {
	bq.fields = append(bq.fields, fields...)
	selbuild := &BlobSelect{BlobQuery: bq}
	selbuild.label = blob.Label
	selbuild.flds, selbuild.scan = &bq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a BlobSelect configured with the given aggregations.
func (bq *BlobQuery) Aggregate(fns ...AggregateFunc) *BlobSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BlobQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bq.fields {
		if !blob.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BlobQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Blob, error) {
	var (
		nodes       = []*Blob{}
		_spec       = bq.querySpec()
		loadedTypes = [2]bool{
			bq.withAttachments != nil,
			bq.withImages != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Blob).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Blob{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withAttachments; query != nil {
		if err := bq.loadAttachments(ctx, query, nodes,
			func(n *Blob) { n.Edges.Attachments = []*Attachment{} },
			func(n *Blob, e *Attachment) { n.Edges.Attachments = append(n.Edges.Attachments, e) }); err != nil {
			return nil, err
		}
	}
	if query := bq.withImages; query != nil {
		if err := bq.loadImages(ctx, query, nodes,
			func(n *Blob) { n.Edges.Images = []*Image{} },
			func(n *Blob, e *Image) { n.Edges.Images = append(n.Edges.Images, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BlobQuery) loadAttachments(ctx context.Context, query *AttachmentQuery, nodes []*Blob, init func(*Blob), assign func(*Blob, *Attachment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Blob)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.InValues(blob.AttachmentsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.blob_attachments
		if fk == nil {
			return fmt.Errorf(`foreign-key "blob_attachments" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "blob_attachments" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (bq *BlobQuery) loadImages(ctx context.Context, query *ImageQuery, nodes []*Blob, init func(*Blob), assign func(*Blob, *Image)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Blob)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Image(func(s *sql.Selector) {
		s.Where(sql.InValues(blob.ImagesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BlobId
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "blobId" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BlobQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.fields
	if len(bq.fields) > 0 {
		_spec.Unique = bq.unique != nil && *bq.unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BlobQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (bq *BlobQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blob.Table,
			Columns: blob.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		},
		From:   bq.sql,
		Unique: true,
	}
	if unique := bq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blob.FieldID)
		for i := range fields {
			if fields[i] != blob.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BlobQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(blob.Table)
	columns := bq.fields
	if len(columns) == 0 {
		columns = blob.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.unique != nil && *bq.unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BlobGroupBy is the group-by builder for Blob entities.
type BlobGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BlobGroupBy) Aggregate(fns ...AggregateFunc) *BlobGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bgb *BlobGroupBy) Scan(ctx context.Context, v any) error {
	query, err := bgb.path(ctx)
	if err != nil {
		return err
	}
	bgb.sql = query
	return bgb.sqlScan(ctx, v)
}

func (bgb *BlobGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range bgb.fields {
		if !blob.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BlobGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql.Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
		for _, f := range bgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bgb.fields...)...)
}

// BlobSelect is the builder for selecting fields of Blob entities.
type BlobSelect struct {
	*BlobQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BlobSelect) Aggregate(fns ...AggregateFunc) *BlobSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BlobSelect) Scan(ctx context.Context, v any) error {
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	bs.sql = bs.BlobQuery.sqlQuery(ctx)
	return bs.sqlScan(ctx, v)
}

func (bs *BlobSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(bs.sql))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		bs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		bs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := bs.sql.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

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
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/ent/gallery"
	"github.com/kzmijak/zswod_api_go/ent/image"
	"github.com/kzmijak/zswod_api_go/ent/predicate"
	"github.com/kzmijak/zswod_api_go/ent/user"
)

// GalleryQuery is the builder for querying Gallery entities.
type GalleryQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.Gallery
	withImages  *ImageQuery
	withArticle *ArticleQuery
	withAuthor  *UserQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GalleryQuery builder.
func (gq *GalleryQuery) Where(ps ...predicate.Gallery) *GalleryQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GalleryQuery) Limit(limit int) *GalleryQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GalleryQuery) Offset(offset int) *GalleryQuery {
	gq.offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GalleryQuery) Unique(unique bool) *GalleryQuery {
	gq.unique = &unique
	return gq
}

// Order adds an order step to the query.
func (gq *GalleryQuery) Order(o ...OrderFunc) *GalleryQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryImages chains the current query on the "images" edge.
func (gq *GalleryQuery) QueryImages() *ImageQuery {
	query := &ImageQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gallery.Table, gallery.FieldID, selector),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, gallery.ImagesTable, gallery.ImagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryArticle chains the current query on the "article" edge.
func (gq *GalleryQuery) QueryArticle() *ArticleQuery {
	query := &ArticleQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gallery.Table, gallery.FieldID, selector),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, gallery.ArticleTable, gallery.ArticleColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAuthor chains the current query on the "author" edge.
func (gq *GalleryQuery) QueryAuthor() *UserQuery {
	query := &UserQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gallery.Table, gallery.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, gallery.AuthorTable, gallery.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Gallery entity from the query.
// Returns a *NotFoundError when no Gallery was found.
func (gq *GalleryQuery) First(ctx context.Context) (*Gallery, error) {
	nodes, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gallery.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GalleryQuery) FirstX(ctx context.Context) *Gallery {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Gallery ID from the query.
// Returns a *NotFoundError when no Gallery ID was found.
func (gq *GalleryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gallery.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GalleryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Gallery entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Gallery entity is found.
// Returns a *NotFoundError when no Gallery entities are found.
func (gq *GalleryQuery) Only(ctx context.Context) (*Gallery, error) {
	nodes, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gallery.Label}
	default:
		return nil, &NotSingularError{gallery.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GalleryQuery) OnlyX(ctx context.Context) *Gallery {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Gallery ID in the query.
// Returns a *NotSingularError when more than one Gallery ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GalleryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gallery.Label}
	default:
		err = &NotSingularError{gallery.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GalleryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Galleries.
func (gq *GalleryQuery) All(ctx context.Context) ([]*Gallery, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gq *GalleryQuery) AllX(ctx context.Context) []*Gallery {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Gallery IDs.
func (gq *GalleryQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := gq.Select(gallery.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GalleryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GalleryQuery) Count(ctx context.Context) (int, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GalleryQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GalleryQuery) Exist(ctx context.Context) (bool, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GalleryQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GalleryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GalleryQuery) Clone() *GalleryQuery {
	if gq == nil {
		return nil
	}
	return &GalleryQuery{
		config:      gq.config,
		limit:       gq.limit,
		offset:      gq.offset,
		order:       append([]OrderFunc{}, gq.order...),
		predicates:  append([]predicate.Gallery{}, gq.predicates...),
		withImages:  gq.withImages.Clone(),
		withArticle: gq.withArticle.Clone(),
		withAuthor:  gq.withAuthor.Clone(),
		// clone intermediate query.
		sql:    gq.sql.Clone(),
		path:   gq.path,
		unique: gq.unique,
	}
}

// WithImages tells the query-builder to eager-load the nodes that are connected to
// the "images" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GalleryQuery) WithImages(opts ...func(*ImageQuery)) *GalleryQuery {
	query := &ImageQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withImages = query
	return gq
}

// WithArticle tells the query-builder to eager-load the nodes that are connected to
// the "article" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GalleryQuery) WithArticle(opts ...func(*ArticleQuery)) *GalleryQuery {
	query := &ArticleQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withArticle = query
	return gq
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GalleryQuery) WithAuthor(opts ...func(*UserQuery)) *GalleryQuery {
	query := &UserQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withAuthor = query
	return gq
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
//	client.Gallery.Query().
//		GroupBy(gallery.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gq *GalleryQuery) GroupBy(field string, fields ...string) *GalleryGroupBy {
	grbuild := &GalleryGroupBy{config: gq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gq.sqlQuery(ctx), nil
	}
	grbuild.label = gallery.Label
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
//	client.Gallery.Query().
//		Select(gallery.FieldCreateTime).
//		Scan(ctx, &v)
func (gq *GalleryQuery) Select(fields ...string) *GallerySelect {
	gq.fields = append(gq.fields, fields...)
	selbuild := &GallerySelect{GalleryQuery: gq}
	selbuild.label = gallery.Label
	selbuild.flds, selbuild.scan = &gq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a GallerySelect configured with the given aggregations.
func (gq *GalleryQuery) Aggregate(fns ...AggregateFunc) *GallerySelect {
	return gq.Select().Aggregate(fns...)
}

func (gq *GalleryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gq.fields {
		if !gallery.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GalleryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Gallery, error) {
	var (
		nodes       = []*Gallery{}
		withFKs     = gq.withFKs
		_spec       = gq.querySpec()
		loadedTypes = [3]bool{
			gq.withImages != nil,
			gq.withArticle != nil,
			gq.withAuthor != nil,
		}
	)
	if gq.withArticle != nil || gq.withAuthor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, gallery.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Gallery).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Gallery{config: gq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gq.withImages; query != nil {
		if err := gq.loadImages(ctx, query, nodes,
			func(n *Gallery) { n.Edges.Images = []*Image{} },
			func(n *Gallery, e *Image) { n.Edges.Images = append(n.Edges.Images, e) }); err != nil {
			return nil, err
		}
	}
	if query := gq.withArticle; query != nil {
		if err := gq.loadArticle(ctx, query, nodes, nil,
			func(n *Gallery, e *Article) { n.Edges.Article = e }); err != nil {
			return nil, err
		}
	}
	if query := gq.withAuthor; query != nil {
		if err := gq.loadAuthor(ctx, query, nodes, nil,
			func(n *Gallery, e *User) { n.Edges.Author = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gq *GalleryQuery) loadImages(ctx context.Context, query *ImageQuery, nodes []*Gallery, init func(*Gallery), assign func(*Gallery, *Image)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Gallery)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Image(func(s *sql.Selector) {
		s.Where(sql.InValues(gallery.ImagesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.gallery_images
		if fk == nil {
			return fmt.Errorf(`foreign-key "gallery_images" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "gallery_images" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (gq *GalleryQuery) loadArticle(ctx context.Context, query *ArticleQuery, nodes []*Gallery, init func(*Gallery), assign func(*Gallery, *Article)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Gallery)
	for i := range nodes {
		if nodes[i].article_gallery == nil {
			continue
		}
		fk := *nodes[i].article_gallery
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
			return fmt.Errorf(`unexpected foreign-key "article_gallery" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gq *GalleryQuery) loadAuthor(ctx context.Context, query *UserQuery, nodes []*Gallery, init func(*Gallery), assign func(*Gallery, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Gallery)
	for i := range nodes {
		if nodes[i].user_galleries == nil {
			continue
		}
		fk := *nodes[i].user_galleries
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_galleries" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gq *GalleryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	_spec.Node.Columns = gq.fields
	if len(gq.fields) > 0 {
		_spec.Unique = gq.unique != nil && *gq.unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GalleryQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := gq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (gq *GalleryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gallery.Table,
			Columns: gallery.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: gallery.FieldID,
			},
		},
		From:   gq.sql,
		Unique: true,
	}
	if unique := gq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gallery.FieldID)
		for i := range fields {
			if fields[i] != gallery.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GalleryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(gallery.Table)
	columns := gq.fields
	if len(columns) == 0 {
		columns = gallery.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.unique != nil && *gq.unique {
		selector.Distinct()
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GalleryGroupBy is the group-by builder for Gallery entities.
type GalleryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GalleryGroupBy) Aggregate(fns ...AggregateFunc) *GalleryGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the group-by query and scans the result into the given value.
func (ggb *GalleryGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ggb.path(ctx)
	if err != nil {
		return err
	}
	ggb.sql = query
	return ggb.sqlScan(ctx, v)
}

func (ggb *GalleryGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ggb.fields {
		if !gallery.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ggb *GalleryGroupBy) sqlQuery() *sql.Selector {
	selector := ggb.sql.Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ggb.fields)+len(ggb.fns))
		for _, f := range ggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ggb.fields...)...)
}

// GallerySelect is the builder for selecting fields of Gallery entities.
type GallerySelect struct {
	*GalleryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GallerySelect) Aggregate(fns ...AggregateFunc) *GallerySelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GallerySelect) Scan(ctx context.Context, v any) error {
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	gs.sql = gs.GalleryQuery.sqlQuery(ctx)
	return gs.sqlScan(ctx, v)
}

func (gs *GallerySelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(gs.sql))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		gs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		gs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := gs.sql.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

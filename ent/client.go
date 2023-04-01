// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/migrate"

	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/ent/attachment"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/ent/custompage"
	"github.com/kzmijak/zswod_api_go/ent/gallery"
	"github.com/kzmijak/zswod_api_go/ent/image"
	"github.com/kzmijak/zswod_api_go/ent/resetpasswordtoken"
	"github.com/kzmijak/zswod_api_go/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Article is the client for interacting with the Article builders.
	Article *ArticleClient
	// Attachment is the client for interacting with the Attachment builders.
	Attachment *AttachmentClient
	// Blob is the client for interacting with the Blob builders.
	Blob *BlobClient
	// CustomPage is the client for interacting with the CustomPage builders.
	CustomPage *CustomPageClient
	// Gallery is the client for interacting with the Gallery builders.
	Gallery *GalleryClient
	// Image is the client for interacting with the Image builders.
	Image *ImageClient
	// ResetPasswordToken is the client for interacting with the ResetPasswordToken builders.
	ResetPasswordToken *ResetPasswordTokenClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Article = NewArticleClient(c.config)
	c.Attachment = NewAttachmentClient(c.config)
	c.Blob = NewBlobClient(c.config)
	c.CustomPage = NewCustomPageClient(c.config)
	c.Gallery = NewGalleryClient(c.config)
	c.Image = NewImageClient(c.config)
	c.ResetPasswordToken = NewResetPasswordTokenClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                ctx,
		config:             cfg,
		Article:            NewArticleClient(cfg),
		Attachment:         NewAttachmentClient(cfg),
		Blob:               NewBlobClient(cfg),
		CustomPage:         NewCustomPageClient(cfg),
		Gallery:            NewGalleryClient(cfg),
		Image:              NewImageClient(cfg),
		ResetPasswordToken: NewResetPasswordTokenClient(cfg),
		User:               NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                ctx,
		config:             cfg,
		Article:            NewArticleClient(cfg),
		Attachment:         NewAttachmentClient(cfg),
		Blob:               NewBlobClient(cfg),
		CustomPage:         NewCustomPageClient(cfg),
		Gallery:            NewGalleryClient(cfg),
		Image:              NewImageClient(cfg),
		ResetPasswordToken: NewResetPasswordTokenClient(cfg),
		User:               NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Article.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Article.Use(hooks...)
	c.Attachment.Use(hooks...)
	c.Blob.Use(hooks...)
	c.CustomPage.Use(hooks...)
	c.Gallery.Use(hooks...)
	c.Image.Use(hooks...)
	c.ResetPasswordToken.Use(hooks...)
	c.User.Use(hooks...)
}

// ArticleClient is a client for the Article schema.
type ArticleClient struct {
	config
}

// NewArticleClient returns a client for the Article from the given config.
func NewArticleClient(c config) *ArticleClient {
	return &ArticleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `article.Hooks(f(g(h())))`.
func (c *ArticleClient) Use(hooks ...Hook) {
	c.hooks.Article = append(c.hooks.Article, hooks...)
}

// Create returns a builder for creating a Article entity.
func (c *ArticleClient) Create() *ArticleCreate {
	mutation := newArticleMutation(c.config, OpCreate)
	return &ArticleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Article entities.
func (c *ArticleClient) CreateBulk(builders ...*ArticleCreate) *ArticleCreateBulk {
	return &ArticleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Article.
func (c *ArticleClient) Update() *ArticleUpdate {
	mutation := newArticleMutation(c.config, OpUpdate)
	return &ArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArticleClient) UpdateOne(a *Article) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticle(a))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArticleClient) UpdateOneID(id uuid.UUID) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticleID(id))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Article.
func (c *ArticleClient) Delete() *ArticleDelete {
	mutation := newArticleMutation(c.config, OpDelete)
	return &ArticleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ArticleClient) DeleteOne(a *Article) *ArticleDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ArticleClient) DeleteOneID(id uuid.UUID) *ArticleDeleteOne {
	builder := c.Delete().Where(article.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArticleDeleteOne{builder}
}

// Query returns a query builder for Article.
func (c *ArticleClient) Query() *ArticleQuery {
	return &ArticleQuery{
		config: c.config,
	}
}

// Get returns a Article entity by its id.
func (c *ArticleClient) Get(ctx context.Context, id uuid.UUID) (*Article, error) {
	return c.Query().Where(article.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArticleClient) GetX(ctx context.Context, id uuid.UUID) *Article {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGallery queries the gallery edge of a Article.
func (c *ArticleClient) QueryGallery(a *Article) *GalleryQuery {
	query := &GalleryQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, id),
			sqlgraph.To(gallery.Table, gallery.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, article.GalleryTable, article.GalleryColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthor queries the author edge of a Article.
func (c *ArticleClient) QueryAuthor(a *Article) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, article.AuthorTable, article.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAttachments queries the attachments edge of a Article.
func (c *ArticleClient) QueryAttachments(a *Article) *AttachmentQuery {
	query := &AttachmentQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, id),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, article.AttachmentsTable, article.AttachmentsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArticleClient) Hooks() []Hook {
	return c.hooks.Article
}

// AttachmentClient is a client for the Attachment schema.
type AttachmentClient struct {
	config
}

// NewAttachmentClient returns a client for the Attachment from the given config.
func NewAttachmentClient(c config) *AttachmentClient {
	return &AttachmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `attachment.Hooks(f(g(h())))`.
func (c *AttachmentClient) Use(hooks ...Hook) {
	c.hooks.Attachment = append(c.hooks.Attachment, hooks...)
}

// Create returns a builder for creating a Attachment entity.
func (c *AttachmentClient) Create() *AttachmentCreate {
	mutation := newAttachmentMutation(c.config, OpCreate)
	return &AttachmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Attachment entities.
func (c *AttachmentClient) CreateBulk(builders ...*AttachmentCreate) *AttachmentCreateBulk {
	return &AttachmentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Attachment.
func (c *AttachmentClient) Update() *AttachmentUpdate {
	mutation := newAttachmentMutation(c.config, OpUpdate)
	return &AttachmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AttachmentClient) UpdateOne(a *Attachment) *AttachmentUpdateOne {
	mutation := newAttachmentMutation(c.config, OpUpdateOne, withAttachment(a))
	return &AttachmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AttachmentClient) UpdateOneID(id uuid.UUID) *AttachmentUpdateOne {
	mutation := newAttachmentMutation(c.config, OpUpdateOne, withAttachmentID(id))
	return &AttachmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Attachment.
func (c *AttachmentClient) Delete() *AttachmentDelete {
	mutation := newAttachmentMutation(c.config, OpDelete)
	return &AttachmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AttachmentClient) DeleteOne(a *Attachment) *AttachmentDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AttachmentClient) DeleteOneID(id uuid.UUID) *AttachmentDeleteOne {
	builder := c.Delete().Where(attachment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AttachmentDeleteOne{builder}
}

// Query returns a query builder for Attachment.
func (c *AttachmentClient) Query() *AttachmentQuery {
	return &AttachmentQuery{
		config: c.config,
	}
}

// Get returns a Attachment entity by its id.
func (c *AttachmentClient) Get(ctx context.Context, id uuid.UUID) (*Attachment, error) {
	return c.Query().Where(attachment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AttachmentClient) GetX(ctx context.Context, id uuid.UUID) *Attachment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBlob queries the blob edge of a Attachment.
func (c *AttachmentClient) QueryBlob(a *Attachment) *BlobQuery {
	query := &BlobQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, id),
			sqlgraph.To(blob.Table, blob.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, attachment.BlobTable, attachment.BlobColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AttachmentClient) Hooks() []Hook {
	return c.hooks.Attachment
}

// BlobClient is a client for the Blob schema.
type BlobClient struct {
	config
}

// NewBlobClient returns a client for the Blob from the given config.
func NewBlobClient(c config) *BlobClient {
	return &BlobClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `blob.Hooks(f(g(h())))`.
func (c *BlobClient) Use(hooks ...Hook) {
	c.hooks.Blob = append(c.hooks.Blob, hooks...)
}

// Create returns a builder for creating a Blob entity.
func (c *BlobClient) Create() *BlobCreate {
	mutation := newBlobMutation(c.config, OpCreate)
	return &BlobCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Blob entities.
func (c *BlobClient) CreateBulk(builders ...*BlobCreate) *BlobCreateBulk {
	return &BlobCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Blob.
func (c *BlobClient) Update() *BlobUpdate {
	mutation := newBlobMutation(c.config, OpUpdate)
	return &BlobUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BlobClient) UpdateOne(b *Blob) *BlobUpdateOne {
	mutation := newBlobMutation(c.config, OpUpdateOne, withBlob(b))
	return &BlobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BlobClient) UpdateOneID(id uuid.UUID) *BlobUpdateOne {
	mutation := newBlobMutation(c.config, OpUpdateOne, withBlobID(id))
	return &BlobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Blob.
func (c *BlobClient) Delete() *BlobDelete {
	mutation := newBlobMutation(c.config, OpDelete)
	return &BlobDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BlobClient) DeleteOne(b *Blob) *BlobDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BlobClient) DeleteOneID(id uuid.UUID) *BlobDeleteOne {
	builder := c.Delete().Where(blob.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BlobDeleteOne{builder}
}

// Query returns a query builder for Blob.
func (c *BlobClient) Query() *BlobQuery {
	return &BlobQuery{
		config: c.config,
	}
}

// Get returns a Blob entity by its id.
func (c *BlobClient) Get(ctx context.Context, id uuid.UUID) (*Blob, error) {
	return c.Query().Where(blob.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BlobClient) GetX(ctx context.Context, id uuid.UUID) *Blob {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BlobClient) Hooks() []Hook {
	return c.hooks.Blob
}

// CustomPageClient is a client for the CustomPage schema.
type CustomPageClient struct {
	config
}

// NewCustomPageClient returns a client for the CustomPage from the given config.
func NewCustomPageClient(c config) *CustomPageClient {
	return &CustomPageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `custompage.Hooks(f(g(h())))`.
func (c *CustomPageClient) Use(hooks ...Hook) {
	c.hooks.CustomPage = append(c.hooks.CustomPage, hooks...)
}

// Create returns a builder for creating a CustomPage entity.
func (c *CustomPageClient) Create() *CustomPageCreate {
	mutation := newCustomPageMutation(c.config, OpCreate)
	return &CustomPageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CustomPage entities.
func (c *CustomPageClient) CreateBulk(builders ...*CustomPageCreate) *CustomPageCreateBulk {
	return &CustomPageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CustomPage.
func (c *CustomPageClient) Update() *CustomPageUpdate {
	mutation := newCustomPageMutation(c.config, OpUpdate)
	return &CustomPageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomPageClient) UpdateOne(cp *CustomPage) *CustomPageUpdateOne {
	mutation := newCustomPageMutation(c.config, OpUpdateOne, withCustomPage(cp))
	return &CustomPageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomPageClient) UpdateOneID(id uuid.UUID) *CustomPageUpdateOne {
	mutation := newCustomPageMutation(c.config, OpUpdateOne, withCustomPageID(id))
	return &CustomPageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CustomPage.
func (c *CustomPageClient) Delete() *CustomPageDelete {
	mutation := newCustomPageMutation(c.config, OpDelete)
	return &CustomPageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CustomPageClient) DeleteOne(cp *CustomPage) *CustomPageDeleteOne {
	return c.DeleteOneID(cp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CustomPageClient) DeleteOneID(id uuid.UUID) *CustomPageDeleteOne {
	builder := c.Delete().Where(custompage.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomPageDeleteOne{builder}
}

// Query returns a query builder for CustomPage.
func (c *CustomPageClient) Query() *CustomPageQuery {
	return &CustomPageQuery{
		config: c.config,
	}
}

// Get returns a CustomPage entity by its id.
func (c *CustomPageClient) Get(ctx context.Context, id uuid.UUID) (*CustomPage, error) {
	return c.Query().Where(custompage.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomPageClient) GetX(ctx context.Context, id uuid.UUID) *CustomPage {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAttachments queries the attachments edge of a CustomPage.
func (c *CustomPageClient) QueryAttachments(cp *CustomPage) *AttachmentQuery {
	query := &AttachmentQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(custompage.Table, custompage.FieldID, id),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, custompage.AttachmentsTable, custompage.AttachmentsColumn),
		)
		fromV = sqlgraph.Neighbors(cp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CustomPageClient) Hooks() []Hook {
	return c.hooks.CustomPage
}

// GalleryClient is a client for the Gallery schema.
type GalleryClient struct {
	config
}

// NewGalleryClient returns a client for the Gallery from the given config.
func NewGalleryClient(c config) *GalleryClient {
	return &GalleryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `gallery.Hooks(f(g(h())))`.
func (c *GalleryClient) Use(hooks ...Hook) {
	c.hooks.Gallery = append(c.hooks.Gallery, hooks...)
}

// Create returns a builder for creating a Gallery entity.
func (c *GalleryClient) Create() *GalleryCreate {
	mutation := newGalleryMutation(c.config, OpCreate)
	return &GalleryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Gallery entities.
func (c *GalleryClient) CreateBulk(builders ...*GalleryCreate) *GalleryCreateBulk {
	return &GalleryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Gallery.
func (c *GalleryClient) Update() *GalleryUpdate {
	mutation := newGalleryMutation(c.config, OpUpdate)
	return &GalleryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GalleryClient) UpdateOne(ga *Gallery) *GalleryUpdateOne {
	mutation := newGalleryMutation(c.config, OpUpdateOne, withGallery(ga))
	return &GalleryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GalleryClient) UpdateOneID(id uuid.UUID) *GalleryUpdateOne {
	mutation := newGalleryMutation(c.config, OpUpdateOne, withGalleryID(id))
	return &GalleryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Gallery.
func (c *GalleryClient) Delete() *GalleryDelete {
	mutation := newGalleryMutation(c.config, OpDelete)
	return &GalleryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GalleryClient) DeleteOne(ga *Gallery) *GalleryDeleteOne {
	return c.DeleteOneID(ga.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GalleryClient) DeleteOneID(id uuid.UUID) *GalleryDeleteOne {
	builder := c.Delete().Where(gallery.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GalleryDeleteOne{builder}
}

// Query returns a query builder for Gallery.
func (c *GalleryClient) Query() *GalleryQuery {
	return &GalleryQuery{
		config: c.config,
	}
}

// Get returns a Gallery entity by its id.
func (c *GalleryClient) Get(ctx context.Context, id uuid.UUID) (*Gallery, error) {
	return c.Query().Where(gallery.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GalleryClient) GetX(ctx context.Context, id uuid.UUID) *Gallery {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryImages queries the images edge of a Gallery.
func (c *GalleryClient) QueryImages(ga *Gallery) *ImageQuery {
	query := &ImageQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(gallery.Table, gallery.FieldID, id),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, gallery.ImagesTable, gallery.ImagesColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryArticle queries the article edge of a Gallery.
func (c *GalleryClient) QueryArticle(ga *Gallery) *ArticleQuery {
	query := &ArticleQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(gallery.Table, gallery.FieldID, id),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, gallery.ArticleTable, gallery.ArticleColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthor queries the author edge of a Gallery.
func (c *GalleryClient) QueryAuthor(ga *Gallery) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(gallery.Table, gallery.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, gallery.AuthorTable, gallery.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GalleryClient) Hooks() []Hook {
	return c.hooks.Gallery
}

// ImageClient is a client for the Image schema.
type ImageClient struct {
	config
}

// NewImageClient returns a client for the Image from the given config.
func NewImageClient(c config) *ImageClient {
	return &ImageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `image.Hooks(f(g(h())))`.
func (c *ImageClient) Use(hooks ...Hook) {
	c.hooks.Image = append(c.hooks.Image, hooks...)
}

// Create returns a builder for creating a Image entity.
func (c *ImageClient) Create() *ImageCreate {
	mutation := newImageMutation(c.config, OpCreate)
	return &ImageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Image entities.
func (c *ImageClient) CreateBulk(builders ...*ImageCreate) *ImageCreateBulk {
	return &ImageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Image.
func (c *ImageClient) Update() *ImageUpdate {
	mutation := newImageMutation(c.config, OpUpdate)
	return &ImageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImageClient) UpdateOne(i *Image) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImage(i))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImageClient) UpdateOneID(id uuid.UUID) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImageID(id))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Image.
func (c *ImageClient) Delete() *ImageDelete {
	mutation := newImageMutation(c.config, OpDelete)
	return &ImageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ImageClient) DeleteOne(i *Image) *ImageDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ImageClient) DeleteOneID(id uuid.UUID) *ImageDeleteOne {
	builder := c.Delete().Where(image.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImageDeleteOne{builder}
}

// Query returns a query builder for Image.
func (c *ImageClient) Query() *ImageQuery {
	return &ImageQuery{
		config: c.config,
	}
}

// Get returns a Image entity by its id.
func (c *ImageClient) Get(ctx context.Context, id uuid.UUID) (*Image, error) {
	return c.Query().Where(image.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImageClient) GetX(ctx context.Context, id uuid.UUID) *Image {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGallery queries the gallery edge of a Image.
func (c *ImageClient) QueryGallery(i *Image) *GalleryQuery {
	query := &GalleryQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, id),
			sqlgraph.To(gallery.Table, gallery.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, image.GalleryTable, image.GalleryColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ImageClient) Hooks() []Hook {
	return c.hooks.Image
}

// ResetPasswordTokenClient is a client for the ResetPasswordToken schema.
type ResetPasswordTokenClient struct {
	config
}

// NewResetPasswordTokenClient returns a client for the ResetPasswordToken from the given config.
func NewResetPasswordTokenClient(c config) *ResetPasswordTokenClient {
	return &ResetPasswordTokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `resetpasswordtoken.Hooks(f(g(h())))`.
func (c *ResetPasswordTokenClient) Use(hooks ...Hook) {
	c.hooks.ResetPasswordToken = append(c.hooks.ResetPasswordToken, hooks...)
}

// Create returns a builder for creating a ResetPasswordToken entity.
func (c *ResetPasswordTokenClient) Create() *ResetPasswordTokenCreate {
	mutation := newResetPasswordTokenMutation(c.config, OpCreate)
	return &ResetPasswordTokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ResetPasswordToken entities.
func (c *ResetPasswordTokenClient) CreateBulk(builders ...*ResetPasswordTokenCreate) *ResetPasswordTokenCreateBulk {
	return &ResetPasswordTokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ResetPasswordToken.
func (c *ResetPasswordTokenClient) Update() *ResetPasswordTokenUpdate {
	mutation := newResetPasswordTokenMutation(c.config, OpUpdate)
	return &ResetPasswordTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ResetPasswordTokenClient) UpdateOne(rpt *ResetPasswordToken) *ResetPasswordTokenUpdateOne {
	mutation := newResetPasswordTokenMutation(c.config, OpUpdateOne, withResetPasswordToken(rpt))
	return &ResetPasswordTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ResetPasswordTokenClient) UpdateOneID(id uuid.UUID) *ResetPasswordTokenUpdateOne {
	mutation := newResetPasswordTokenMutation(c.config, OpUpdateOne, withResetPasswordTokenID(id))
	return &ResetPasswordTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ResetPasswordToken.
func (c *ResetPasswordTokenClient) Delete() *ResetPasswordTokenDelete {
	mutation := newResetPasswordTokenMutation(c.config, OpDelete)
	return &ResetPasswordTokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ResetPasswordTokenClient) DeleteOne(rpt *ResetPasswordToken) *ResetPasswordTokenDeleteOne {
	return c.DeleteOneID(rpt.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ResetPasswordTokenClient) DeleteOneID(id uuid.UUID) *ResetPasswordTokenDeleteOne {
	builder := c.Delete().Where(resetpasswordtoken.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ResetPasswordTokenDeleteOne{builder}
}

// Query returns a query builder for ResetPasswordToken.
func (c *ResetPasswordTokenClient) Query() *ResetPasswordTokenQuery {
	return &ResetPasswordTokenQuery{
		config: c.config,
	}
}

// Get returns a ResetPasswordToken entity by its id.
func (c *ResetPasswordTokenClient) Get(ctx context.Context, id uuid.UUID) (*ResetPasswordToken, error) {
	return c.Query().Where(resetpasswordtoken.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ResetPasswordTokenClient) GetX(ctx context.Context, id uuid.UUID) *ResetPasswordToken {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a ResetPasswordToken.
func (c *ResetPasswordTokenClient) QueryOwner(rpt *ResetPasswordToken) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rpt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(resetpasswordtoken.Table, resetpasswordtoken.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, resetpasswordtoken.OwnerTable, resetpasswordtoken.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(rpt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ResetPasswordTokenClient) Hooks() []Hook {
	return c.hooks.ResetPasswordToken
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGalleries queries the galleries edge of a User.
func (c *UserClient) QueryGalleries(u *User) *GalleryQuery {
	query := &GalleryQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(gallery.Table, gallery.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.GalleriesTable, user.GalleriesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryArticles queries the articles edge of a User.
func (c *UserClient) QueryArticles(u *User) *ArticleQuery {
	query := &ArticleQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ArticlesTable, user.ArticlesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAvatar queries the avatar edge of a User.
func (c *UserClient) QueryAvatar(u *User) *ImageQuery {
	query := &ImageQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, user.AvatarTable, user.AvatarColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryResetPasswordToken queries the resetPasswordToken edge of a User.
func (c *UserClient) QueryResetPasswordToken(u *User) *ResetPasswordTokenQuery {
	query := &ResetPasswordTokenQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(resetpasswordtoken.Table, resetpasswordtoken.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.ResetPasswordTokenTable, user.ResetPasswordTokenColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

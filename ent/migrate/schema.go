// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 200},
		{Name: "title_normalized", Type: field.TypeString, Unique: true},
		{Name: "short", Type: field.TypeString, Size: 300},
		{Name: "content", Type: field.TypeString, SchemaType: map[string]string{"mysql": "mediumtext"}},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
	}
	// AttachmentsColumns holds the columns for the "attachments" table.
	AttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "icon_id", Type: field.TypeString, Nullable: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "article_attachments", Type: field.TypeUUID, Nullable: true},
		{Name: "attachment_blob", Type: field.TypeUUID},
		{Name: "custom_page_attachments", Type: field.TypeUUID, Nullable: true},
	}
	// AttachmentsTable holds the schema information for the "attachments" table.
	AttachmentsTable = &schema.Table{
		Name:       "attachments",
		Columns:    AttachmentsColumns,
		PrimaryKey: []*schema.Column{AttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attachments_articles_attachments",
				Columns:    []*schema.Column{AttachmentsColumns[4]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "attachments_blobs_blob",
				Columns:    []*schema.Column{AttachmentsColumns[5]},
				RefColumns: []*schema.Column{BlobsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "attachments_custom_pages_attachments",
				Columns:    []*schema.Column{AttachmentsColumns[6]},
				RefColumns: []*schema.Column{CustomPagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BlobsColumns holds the columns for the "blobs" table.
	BlobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "blob", Type: field.TypeBytes, SchemaType: map[string]string{"mysql": "mediumblob"}},
		{Name: "title", Type: field.TypeString},
		{Name: "alt", Type: field.TypeString},
		{Name: "content_type", Type: field.TypeString},
	}
	// BlobsTable holds the schema information for the "blobs" table.
	BlobsTable = &schema.Table{
		Name:       "blobs",
		Columns:    BlobsColumns,
		PrimaryKey: []*schema.Column{BlobsColumns[0]},
	}
	// CustomPagesColumns holds the columns for the "custom_pages" table.
	CustomPagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "icon_id", Type: field.TypeString, Nullable: true},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 200},
		{Name: "title_normalized", Type: field.TypeString, Unique: true},
		{Name: "content", Type: field.TypeString, SchemaType: map[string]string{"mysql": "mediumtext"}},
		{Name: "is_external", Type: field.TypeBool, Nullable: true},
		{Name: "link", Type: field.TypeString, Nullable: true},
		{Name: "section", Type: field.TypeString},
	}
	// CustomPagesTable holds the schema information for the "custom_pages" table.
	CustomPagesTable = &schema.Table{
		Name:       "custom_pages",
		Columns:    CustomPagesColumns,
		PrimaryKey: []*schema.Column{CustomPagesColumns[0]},
	}
	// GalleriesColumns holds the columns for the "galleries" table.
	GalleriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "article_gallery", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "user_galleries", Type: field.TypeUUID},
	}
	// GalleriesTable holds the schema information for the "galleries" table.
	GalleriesTable = &schema.Table{
		Name:       "galleries",
		Columns:    GalleriesColumns,
		PrimaryKey: []*schema.Column{GalleriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "galleries_articles_gallery",
				Columns:    []*schema.Column{GalleriesColumns[4]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "galleries_users_galleries",
				Columns:    []*schema.Column{GalleriesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "alt", Type: field.TypeString, Nullable: true},
		{Name: "order", Type: field.TypeInt, Nullable: true},
		{Name: "gallery_images", Type: field.TypeUUID, Nullable: true},
		{Name: "blob_id", Type: field.TypeUUID},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "images_galleries_images",
				Columns:    []*schema.Column{ImagesColumns[4]},
				RefColumns: []*schema.Column{GalleriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "images_blobs_blob",
				Columns:    []*schema.Column{ImagesColumns[5]},
				RefColumns: []*schema.Column{BlobsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ResetPasswordTokensColumns holds the columns for the "reset_password_tokens" table.
	ResetPasswordTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_reset_password_tokens", Type: field.TypeUUID},
	}
	// ResetPasswordTokensTable holds the schema information for the "reset_password_tokens" table.
	ResetPasswordTokensTable = &schema.Table{
		Name:       "reset_password_tokens",
		Columns:    ResetPasswordTokensColumns,
		PrimaryKey: []*schema.Column{ResetPasswordTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reset_password_tokens_users_resetPasswordTokens",
				Columns:    []*schema.Column{ResetPasswordTokensColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"Admin", "Teacher", "LegalGuardian", "Student", "Unknown"}},
		{Name: "user_avatar", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_images_avatar",
				Columns:    []*schema.Column{UsersColumns[6]},
				RefColumns: []*schema.Column{ImagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserArticlesColumns holds the columns for the "user_articles" table.
	UserArticlesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "article_id", Type: field.TypeUUID},
	}
	// UserArticlesTable holds the schema information for the "user_articles" table.
	UserArticlesTable = &schema.Table{
		Name:       "user_articles",
		Columns:    UserArticlesColumns,
		PrimaryKey: []*schema.Column{UserArticlesColumns[0], UserArticlesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_articles_user_id",
				Columns:    []*schema.Column{UserArticlesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_articles_article_id",
				Columns:    []*schema.Column{UserArticlesColumns[1]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		AttachmentsTable,
		BlobsTable,
		CustomPagesTable,
		GalleriesTable,
		ImagesTable,
		ResetPasswordTokensTable,
		UsersTable,
		UserArticlesTable,
	}
)

func init() {
	AttachmentsTable.ForeignKeys[0].RefTable = ArticlesTable
	AttachmentsTable.ForeignKeys[1].RefTable = BlobsTable
	AttachmentsTable.ForeignKeys[2].RefTable = CustomPagesTable
	GalleriesTable.ForeignKeys[0].RefTable = ArticlesTable
	GalleriesTable.ForeignKeys[1].RefTable = UsersTable
	ImagesTable.ForeignKeys[0].RefTable = GalleriesTable
	ImagesTable.ForeignKeys[1].RefTable = BlobsTable
	ResetPasswordTokensTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = ImagesTable
	UserArticlesTable.ForeignKeys[0].RefTable = UsersTable
	UserArticlesTable.ForeignKeys[1].RefTable = ArticlesTable
}

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CarouselsColumns holds the columns for the "carousels" table.
	CarouselsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "img_path", Type: field.TypeString},
		{Name: "product_id", Type: field.TypeUint64, Nullable: true},
	}
	// CarouselsTable holds the schema information for the "carousels" table.
	CarouselsTable = &schema.Table{
		Name:       "carousels",
		Columns:    CarouselsColumns,
		PrimaryKey: []*schema.Column{CarouselsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "carousels_products_carousels",
				Columns:    []*schema.Column{CarouselsColumns[4]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "category_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// EmailsColumns holds the columns for the "emails" table.
	EmailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "email", Type: field.TypeString},
		{Name: "is_verified", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "secret", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "user_emails", Type: field.TypeUint64, Nullable: true},
	}
	// EmailsTable holds the schema information for the "emails" table.
	EmailsTable = &schema.Table{
		Name:       "emails",
		Columns:    EmailsColumns,
		PrimaryKey: []*schema.Column{EmailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "emails_users_emails",
				Columns:    []*schema.Column{EmailsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "info", Type: field.TypeString},
		{Name: "price", Type: field.TypeInt64},
		{Name: "discount_price", Type: field.TypeInt64, Nullable: true},
		{Name: "category_id", Type: field.TypeUint64, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_categories_product",
				Columns:    []*schema.Column{ProductsColumns[8]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProductImgsColumns holds the columns for the "product_imgs" table.
	ProductImgsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "img_path", Type: field.TypeString},
		{Name: "product_id", Type: field.TypeUint64, Nullable: true},
	}
	// ProductImgsTable holds the schema information for the "product_imgs" table.
	ProductImgsTable = &schema.Table{
		Name:       "product_imgs",
		Columns:    ProductImgsColumns,
		PrimaryKey: []*schema.Column{ProductImgsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "product_imgs_products_productimgs",
				Columns:    []*schema.Column{ProductImgsColumns[4]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "nickname", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Default: "", SchemaType: map[string]string{"mysql": "varchar(512)"}},
		{Name: "github_id", Type: field.TypeUint64, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CarouselsTable,
		CategoriesTable,
		EmailsTable,
		ProductsTable,
		ProductImgsTable,
		UsersTable,
	}
)

func init() {
	CarouselsTable.ForeignKeys[0].RefTable = ProductsTable
	EmailsTable.ForeignKeys[0].RefTable = UsersTable
	ProductsTable.ForeignKeys[0].RefTable = CategoriesTable
	ProductImgsTable.ForeignKeys[0].RefTable = ProductsTable
}

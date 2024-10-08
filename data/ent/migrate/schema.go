// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AddressesColumns holds the columns for the "addresses" table.
	AddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeUint64, Nullable: true},
	}
	// AddressesTable holds the schema information for the "addresses" table.
	AddressesTable = &schema.Table{
		Name:       "addresses",
		Columns:    AddressesColumns,
		PrimaryKey: []*schema.Column{AddressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "addresses_users_address",
				Columns:    []*schema.Column{AddressesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
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
	// CartsColumns holds the columns for the "carts" table.
	CartsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "num", Type: field.TypeInt32},
		{Name: "max_num", Type: field.TypeInt32},
		{Name: "check", Type: field.TypeBool},
		{Name: "product_id", Type: field.TypeUint64, Nullable: true},
		{Name: "user_id", Type: field.TypeUint64, Nullable: true},
	}
	// CartsTable holds the schema information for the "carts" table.
	CartsTable = &schema.Table{
		Name:       "carts",
		Columns:    CartsColumns,
		PrimaryKey: []*schema.Column{CartsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "carts_products_cart",
				Columns:    []*schema.Column{CartsColumns[6]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "carts_users_cart",
				Columns:    []*schema.Column{CartsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
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
	// FavouritesColumns holds the columns for the "favourites" table.
	FavouritesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "product_id", Type: field.TypeUint64, Nullable: true},
		{Name: "user_id", Type: field.TypeUint64, Nullable: true},
	}
	// FavouritesTable holds the schema information for the "favourites" table.
	FavouritesTable = &schema.Table{
		Name:       "favourites",
		Columns:    FavouritesColumns,
		PrimaryKey: []*schema.Column{FavouritesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "favourites_products_favourite",
				Columns:    []*schema.Column{FavouritesColumns[3]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "favourites_users_favourite",
				Columns:    []*schema.Column{FavouritesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NoticesColumns holds the columns for the "notices" table.
	NoticesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString},
	}
	// NoticesTable holds the schema information for the "notices" table.
	NoticesTable = &schema.Table{
		Name:       "notices",
		Columns:    NoticesColumns,
		PrimaryKey: []*schema.Column{NoticesColumns[0]},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "num", Type: field.TypeInt32},
		{Name: "order_num", Type: field.TypeUint64},
		{Name: "address_name", Type: field.TypeString},
		{Name: "address_phone", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "type", Type: field.TypeUint64},
		{Name: "price", Type: field.TypeInt64},
		{Name: "product_id", Type: field.TypeUint64, Nullable: true},
		{Name: "user_id", Type: field.TypeUint64, Nullable: true},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_products_order",
				Columns:    []*schema.Column{OrdersColumns[10]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "orders_users_order",
				Columns:    []*schema.Column{OrdersColumns[11]},
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
		AddressesTable,
		CarouselsTable,
		CartsTable,
		CategoriesTable,
		EmailsTable,
		FavouritesTable,
		NoticesTable,
		OrdersTable,
		ProductsTable,
		ProductImgsTable,
		UsersTable,
	}
)

func init() {
	AddressesTable.ForeignKeys[0].RefTable = UsersTable
	CarouselsTable.ForeignKeys[0].RefTable = ProductsTable
	CartsTable.ForeignKeys[0].RefTable = ProductsTable
	CartsTable.ForeignKeys[1].RefTable = UsersTable
	EmailsTable.ForeignKeys[0].RefTable = UsersTable
	FavouritesTable.ForeignKeys[0].RefTable = ProductsTable
	FavouritesTable.ForeignKeys[1].RefTable = UsersTable
	OrdersTable.ForeignKeys[0].RefTable = ProductsTable
	OrdersTable.ForeignKeys[1].RefTable = UsersTable
	ProductsTable.ForeignKeys[0].RefTable = CategoriesTable
	ProductImgsTable.ForeignKeys[0].RefTable = ProductsTable
}

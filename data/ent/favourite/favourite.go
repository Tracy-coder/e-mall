// Code generated by ent, DO NOT EDIT.

package favourite

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the favourite type in the database.
	Label = "favourite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldProductID holds the string denoting the productid field in the database.
	FieldProductID = "product_id"
	// FieldUserID holds the string denoting the userid field in the database.
	FieldUserID = "user_id"
	// EdgeProductFavourite holds the string denoting the product_favourite edge name in mutations.
	EdgeProductFavourite = "product_favourite"
	// EdgeUserFavourite holds the string denoting the user_favourite edge name in mutations.
	EdgeUserFavourite = "user_favourite"
	// Table holds the table name of the favourite in the database.
	Table = "favourites"
	// ProductFavouriteTable is the table that holds the product_favourite relation/edge.
	ProductFavouriteTable = "favourites"
	// ProductFavouriteInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductFavouriteInverseTable = "products"
	// ProductFavouriteColumn is the table column denoting the product_favourite relation/edge.
	ProductFavouriteColumn = "product_id"
	// UserFavouriteTable is the table that holds the user_favourite relation/edge.
	UserFavouriteTable = "favourites"
	// UserFavouriteInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserFavouriteInverseTable = "users"
	// UserFavouriteColumn is the table column denoting the user_favourite relation/edge.
	UserFavouriteColumn = "user_id"
)

// Columns holds all SQL columns for favourite fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldProductID,
	FieldUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Favourite queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByProductID orders the results by the productID field.
func ByProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductID, opts...).ToFunc()
}

// ByUserID orders the results by the userID field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByProductFavouriteField orders the results by product_favourite field.
func ByProductFavouriteField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductFavouriteStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserFavouriteField orders the results by user_favourite field.
func ByUserFavouriteField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserFavouriteStep(), sql.OrderByField(field, opts...))
	}
}
func newProductFavouriteStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductFavouriteInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProductFavouriteTable, ProductFavouriteColumn),
	)
}
func newUserFavouriteStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserFavouriteInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserFavouriteTable, UserFavouriteColumn),
	)
}

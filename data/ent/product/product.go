// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCategoryID holds the string denoting the categoryid field in the database.
	FieldCategoryID = "category_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldInfo holds the string denoting the info field in the database.
	FieldInfo = "info"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldDiscountPrice holds the string denoting the discount_price field in the database.
	FieldDiscountPrice = "discount_price"
	// EdgeCarousels holds the string denoting the carousels edge name in mutations.
	EdgeCarousels = "carousels"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeProductimgs holds the string denoting the productimgs edge name in mutations.
	EdgeProductimgs = "productimgs"
	// EdgeFavourite holds the string denoting the favourite edge name in mutations.
	EdgeFavourite = "favourite"
	// Table holds the table name of the product in the database.
	Table = "products"
	// CarouselsTable is the table that holds the carousels relation/edge.
	CarouselsTable = "carousels"
	// CarouselsInverseTable is the table name for the Carousel entity.
	// It exists in this package in order to avoid circular dependency with the "carousel" package.
	CarouselsInverseTable = "carousels"
	// CarouselsColumn is the table column denoting the carousels relation/edge.
	CarouselsColumn = "product_id"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "products"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "category_id"
	// ProductimgsTable is the table that holds the productimgs relation/edge.
	ProductimgsTable = "product_imgs"
	// ProductimgsInverseTable is the table name for the ProductImg entity.
	// It exists in this package in order to avoid circular dependency with the "productimg" package.
	ProductimgsInverseTable = "product_imgs"
	// ProductimgsColumn is the table column denoting the productimgs relation/edge.
	ProductimgsColumn = "product_id"
	// FavouriteTable is the table that holds the favourite relation/edge.
	FavouriteTable = "favourites"
	// FavouriteInverseTable is the table name for the Favourite entity.
	// It exists in this package in order to avoid circular dependency with the "favourite" package.
	FavouriteInverseTable = "favourites"
	// FavouriteColumn is the table column denoting the favourite relation/edge.
	FavouriteColumn = "product_id"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldCategoryID,
	FieldTitle,
	FieldInfo,
	FieldPrice,
	FieldDiscountPrice,
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

// OrderOption defines the ordering options for the Product queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCategoryID orders the results by the categoryID field.
func ByCategoryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByInfo orders the results by the info field.
func ByInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInfo, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByDiscountPrice orders the results by the discount_price field.
func ByDiscountPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDiscountPrice, opts...).ToFunc()
}

// ByCarouselsCount orders the results by carousels count.
func ByCarouselsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCarouselsStep(), opts...)
	}
}

// ByCarousels orders the results by carousels terms.
func ByCarousels(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCarouselsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCategoryField orders the results by category field.
func ByCategoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryStep(), sql.OrderByField(field, opts...))
	}
}

// ByProductimgsCount orders the results by productimgs count.
func ByProductimgsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductimgsStep(), opts...)
	}
}

// ByProductimgs orders the results by productimgs terms.
func ByProductimgs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductimgsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFavouriteCount orders the results by favourite count.
func ByFavouriteCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFavouriteStep(), opts...)
	}
}

// ByFavourite orders the results by favourite terms.
func ByFavourite(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFavouriteStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCarouselsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CarouselsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CarouselsTable, CarouselsColumn),
	)
}
func newCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
	)
}
func newProductimgsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductimgsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductimgsTable, ProductimgsColumn),
	)
}
func newFavouriteStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FavouriteInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FavouriteTable, FavouriteColumn),
	)
}

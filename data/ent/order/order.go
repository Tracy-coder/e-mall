// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUserID holds the string denoting the userid field in the database.
	FieldUserID = "user_id"
	// FieldProductID holds the string denoting the productid field in the database.
	FieldProductID = "product_id"
	// FieldNum holds the string denoting the num field in the database.
	FieldNum = "num"
	// FieldOrderNum holds the string denoting the ordernum field in the database.
	FieldOrderNum = "order_num"
	// FieldAddressName holds the string denoting the addressname field in the database.
	FieldAddressName = "address_name"
	// FieldAddressPhone holds the string denoting the addressphone field in the database.
	FieldAddressPhone = "address_phone"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// EdgeUserOrder holds the string denoting the user_order edge name in mutations.
	EdgeUserOrder = "user_order"
	// EdgeProductOrder holds the string denoting the product_order edge name in mutations.
	EdgeProductOrder = "product_order"
	// Table holds the table name of the order in the database.
	Table = "orders"
	// UserOrderTable is the table that holds the user_order relation/edge.
	UserOrderTable = "orders"
	// UserOrderInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserOrderInverseTable = "users"
	// UserOrderColumn is the table column denoting the user_order relation/edge.
	UserOrderColumn = "user_id"
	// ProductOrderTable is the table that holds the product_order relation/edge.
	ProductOrderTable = "orders"
	// ProductOrderInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductOrderInverseTable = "products"
	// ProductOrderColumn is the table column denoting the product_order relation/edge.
	ProductOrderColumn = "product_id"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUserID,
	FieldProductID,
	FieldNum,
	FieldOrderNum,
	FieldAddressName,
	FieldAddressPhone,
	FieldAddress,
	FieldType,
	FieldPrice,
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

// OrderOption defines the ordering options for the Order queries.
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

// ByUserID orders the results by the UserID field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByProductID orders the results by the ProductID field.
func ByProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductID, opts...).ToFunc()
}

// ByNum orders the results by the Num field.
func ByNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNum, opts...).ToFunc()
}

// ByOrderNum orders the results by the OrderNum field.
func ByOrderNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderNum, opts...).ToFunc()
}

// ByAddressName orders the results by the AddressName field.
func ByAddressName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddressName, opts...).ToFunc()
}

// ByAddressPhone orders the results by the AddressPhone field.
func ByAddressPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddressPhone, opts...).ToFunc()
}

// ByAddress orders the results by the Address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByType orders the results by the Type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByPrice orders the results by the Price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByUserOrderField orders the results by user_order field.
func ByUserOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserOrderStep(), sql.OrderByField(field, opts...))
	}
}

// ByProductOrderField orders the results by product_order field.
func ByProductOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductOrderStep(), sql.OrderByField(field, opts...))
	}
}
func newUserOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserOrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserOrderTable, UserOrderColumn),
	)
}
func newProductOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductOrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProductOrderTable, ProductOrderColumn),
	)
}

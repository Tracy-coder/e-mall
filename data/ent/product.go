// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Tracy-coder/e-mall/data/ent/category"
	"github.com/Tracy-coder/e-mall/data/ent/product"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID uint64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// product name | 商品名称
	Name string `json:"name,omitempty"`
	// category ID | 分类编号
	CategoryID uint64 `json:"categoryID,omitempty"`
	// title | 标题
	Title string `json:"title,omitempty"`
	// info | 详细信息
	Info string `json:"info,omitempty"`
	// price | 商品价格
	Price int64 `json:"price,omitempty"`
	// discount price | 优惠后价格
	DiscountPrice int64 `json:"discount_price,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges        ProductEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Carousels holds the value of the carousels edge.
	Carousels []*Carousel `json:"carousels,omitempty"`
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// Productimgs holds the value of the productimgs edge.
	Productimgs []*ProductImg `json:"productimgs,omitempty"`
	// Favourite holds the value of the favourite edge.
	Favourite []*Favourite `json:"favourite,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// CarouselsOrErr returns the Carousels value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) CarouselsOrErr() ([]*Carousel, error) {
	if e.loadedTypes[0] {
		return e.Carousels, nil
	}
	return nil, &NotLoadedError{edge: "carousels"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) CategoryOrErr() (*Category, error) {
	if e.Category != nil {
		return e.Category, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: category.Label}
	}
	return nil, &NotLoadedError{edge: "category"}
}

// ProductimgsOrErr returns the Productimgs value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ProductimgsOrErr() ([]*ProductImg, error) {
	if e.loadedTypes[2] {
		return e.Productimgs, nil
	}
	return nil, &NotLoadedError{edge: "productimgs"}
}

// FavouriteOrErr returns the Favourite value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) FavouriteOrErr() ([]*Favourite, error) {
	if e.loadedTypes[3] {
		return e.Favourite, nil
	}
	return nil, &NotLoadedError{edge: "favourite"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldID, product.FieldCategoryID, product.FieldPrice, product.FieldDiscountPrice:
			values[i] = new(sql.NullInt64)
		case product.FieldName, product.FieldTitle, product.FieldInfo:
			values[i] = new(sql.NullString)
		case product.FieldCreatedAt, product.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = uint64(value.Int64)
		case product.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case product.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field categoryID", values[i])
			} else if value.Valid {
				pr.CategoryID = uint64(value.Int64)
			}
		case product.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pr.Title = value.String
			}
		case product.FieldInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field info", values[i])
			} else if value.Valid {
				pr.Info = value.String
			}
		case product.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pr.Price = value.Int64
			}
		case product.FieldDiscountPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field discount_price", values[i])
			} else if value.Valid {
				pr.DiscountPrice = value.Int64
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Product.
// This includes values selected through modifiers, order, etc.
func (pr *Product) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryCarousels queries the "carousels" edge of the Product entity.
func (pr *Product) QueryCarousels() *CarouselQuery {
	return NewProductClient(pr.config).QueryCarousels(pr)
}

// QueryCategory queries the "category" edge of the Product entity.
func (pr *Product) QueryCategory() *CategoryQuery {
	return NewProductClient(pr.config).QueryCategory(pr)
}

// QueryProductimgs queries the "productimgs" edge of the Product entity.
func (pr *Product) QueryProductimgs() *ProductImgQuery {
	return NewProductClient(pr.config).QueryProductimgs(pr)
}

// QueryFavourite queries the "favourite" edge of the Product entity.
func (pr *Product) QueryFavourite() *FavouriteQuery {
	return NewProductClient(pr.config).QueryFavourite(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return NewProductClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("categoryID=")
	builder.WriteString(fmt.Sprintf("%v", pr.CategoryID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(pr.Title)
	builder.WriteString(", ")
	builder.WriteString("info=")
	builder.WriteString(pr.Info)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteString(", ")
	builder.WriteString("discount_price=")
	builder.WriteString(fmt.Sprintf("%v", pr.DiscountPrice))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

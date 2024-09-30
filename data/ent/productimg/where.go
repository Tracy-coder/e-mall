// Code generated by ent, DO NOT EDIT.

package productimg

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Tracy-coder/e-mall/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldEQ(FieldUpdatedAt, v))
}

// ImgPath applies equality check predicate on the "img_path" field. It's identical to ImgPathEQ.
func ImgPath(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldEQ(FieldImgPath, v))
}

// ProductID applies equality check predicate on the "productID" field. It's identical to ProductIDEQ.
func ProductID(v uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldEQ(FieldProductID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldLTE(FieldUpdatedAt, v))
}

// ImgPathEQ applies the EQ predicate on the "img_path" field.
func ImgPathEQ(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldEQ(FieldImgPath, v))
}

// ImgPathNEQ applies the NEQ predicate on the "img_path" field.
func ImgPathNEQ(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldNEQ(FieldImgPath, v))
}

// ImgPathIn applies the In predicate on the "img_path" field.
func ImgPathIn(vs ...string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldIn(FieldImgPath, vs...))
}

// ImgPathNotIn applies the NotIn predicate on the "img_path" field.
func ImgPathNotIn(vs ...string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldNotIn(FieldImgPath, vs...))
}

// ImgPathGT applies the GT predicate on the "img_path" field.
func ImgPathGT(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldGT(FieldImgPath, v))
}

// ImgPathGTE applies the GTE predicate on the "img_path" field.
func ImgPathGTE(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldGTE(FieldImgPath, v))
}

// ImgPathLT applies the LT predicate on the "img_path" field.
func ImgPathLT(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldLT(FieldImgPath, v))
}

// ImgPathLTE applies the LTE predicate on the "img_path" field.
func ImgPathLTE(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldLTE(FieldImgPath, v))
}

// ImgPathContains applies the Contains predicate on the "img_path" field.
func ImgPathContains(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldContains(FieldImgPath, v))
}

// ImgPathHasPrefix applies the HasPrefix predicate on the "img_path" field.
func ImgPathHasPrefix(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldHasPrefix(FieldImgPath, v))
}

// ImgPathHasSuffix applies the HasSuffix predicate on the "img_path" field.
func ImgPathHasSuffix(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldHasSuffix(FieldImgPath, v))
}

// ImgPathEqualFold applies the EqualFold predicate on the "img_path" field.
func ImgPathEqualFold(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldEqualFold(FieldImgPath, v))
}

// ImgPathContainsFold applies the ContainsFold predicate on the "img_path" field.
func ImgPathContainsFold(v string) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldContainsFold(FieldImgPath, v))
}

// ProductIDEQ applies the EQ predicate on the "productID" field.
func ProductIDEQ(v uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldEQ(FieldProductID, v))
}

// ProductIDNEQ applies the NEQ predicate on the "productID" field.
func ProductIDNEQ(v uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldNEQ(FieldProductID, v))
}

// ProductIDIn applies the In predicate on the "productID" field.
func ProductIDIn(vs ...uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldIn(FieldProductID, vs...))
}

// ProductIDNotIn applies the NotIn predicate on the "productID" field.
func ProductIDNotIn(vs ...uint64) predicate.ProductImg {
	return predicate.ProductImg(sql.FieldNotIn(FieldProductID, vs...))
}

// ProductIDIsNil applies the IsNil predicate on the "productID" field.
func ProductIDIsNil() predicate.ProductImg {
	return predicate.ProductImg(sql.FieldIsNull(FieldProductID))
}

// ProductIDNotNil applies the NotNil predicate on the "productID" field.
func ProductIDNotNil() predicate.ProductImg {
	return predicate.ProductImg(sql.FieldNotNull(FieldProductID))
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.ProductImg {
	return predicate.ProductImg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.ProductImg {
	return predicate.ProductImg(func(s *sql.Selector) {
		step := newProductStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProductImg) predicate.ProductImg {
	return predicate.ProductImg(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProductImg) predicate.ProductImg {
	return predicate.ProductImg(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProductImg) predicate.ProductImg {
	return predicate.ProductImg(sql.NotPredicates(p))
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/schema/mixins"
)

// Email holds the schema definition for the Email entity.
type ProductImg struct {
	ent.Schema
}

// Fields of the Email.
func (ProductImg) Fields() []ent.Field {
	return []ent.Field{
		field.String("img_path").Comment("image path | 商品图片"),
		field.Uint64("productID").Optional().Comment("product ID | 产品编号"),
	}
}

// Edges of the Email.
// TODO: categoryID 外键
func (ProductImg) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("productimgs").Field("productID").Unique(),
	}
}

func (ProductImg) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

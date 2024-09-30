package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/schema/mixins"
)

// Email holds the schema definition for the Email entity.
type Product struct {
	ent.Schema
}

// Fields of the Email.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("product name | 商品名称"),
		field.Uint64("categoryID").Optional().Comment("category ID | 分类编号"),
		field.String("title").Comment("title | 标题"),
		field.String("info").Comment("info | 详细信息"),
		field.Int64("price").Comment("price | 商品价格"),
		field.Int64("discount_price").Optional().Comment("discount price | 优惠后价格"),
	}
}

// Edges of the Email.
// TODO: categoryID 外键
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("carousels", Carousel.Type),
		edge.From("category", Category.Type).Ref("product").Field("categoryID").Unique(),
		edge.To("productimgs", ProductImg.Type),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

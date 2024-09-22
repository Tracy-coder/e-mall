package schema

import (
	"entgo.io/ent"
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
		field.Uint64("categoryID").Comment("category ID | 分类编号"),
		field.String("title").Comment("title | 标题"),
		field.String("info").Comment("info | 详细信息"),
		field.Int64("price").Comment("price | 商品价格"),
		field.String("img_path").Comment("image path | 商品图片"),
		field.Int64("discount_price").Comment("discount price | 优惠后价格"),
	}
}

// Edges of the Email.
func (Product) Edges() []ent.Edge {
	return nil
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

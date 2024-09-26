package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/schema/mixins"
)

// Email holds the schema definition for the Email entity.
type Carousel struct {
	ent.Schema
}

// Fields of the Email.
func (Carousel) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("ProductID").Optional().Comment("Product ID | 商品编号"),
		field.String("ImgPath").Comment("image path | 图片路径"),
	}
}

// Edges of the Email.
func (Carousel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Product.Type).
			Ref("carousels").
			Unique().Field("ProductID"),
	}
}

func (Carousel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

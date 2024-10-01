package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/schema/mixins"
)

// Email holds the schema definition for the Email entity.
type Favourite struct {
	ent.Schema
}

// Fields of the Email.
func (Favourite) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("productID").Optional().Comment("product ID | 产品编号"),
		field.Uint64("userID").Optional().Comment("user ID | 用户编号"),
	}
}

// Edges of the Email.
func (Favourite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product_favourite", Product.Type).Ref("favourite").Field("productID").Unique(),
		edge.From("user_favourite", User.Type).Ref("favourite").Field("userID").Unique(),
	}
}

func (Favourite) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

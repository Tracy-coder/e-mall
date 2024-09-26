package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Email holds the schema definition for the Email entity.
type Category struct {
	ent.Schema
}

// Fields of the Email.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("category ID,primary key | 分类编号"),
		field.String("categoryName").Comment("category name | 分类标题"),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Comment("created time"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("last update time"),
	}
}

// Edges of the Email.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type),
	}
}

func (Category) Mixin() []ent.Mixin {
	return nil
}

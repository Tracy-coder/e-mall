package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/schema/mixins"
)

// Email holds the schema definition for the Email entity.
type Notice struct {
	ent.Schema
}

// Fields of the Email.
func (Notice) Fields() []ent.Field {
	return []ent.Field{
		field.String("Text").Comment("notice text | 通知详情"),
	}
}

// Edges of the Email.
func (Notice) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Notice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

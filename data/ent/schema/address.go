package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/schema/mixins"
)

// Email holds the schema definition for the Email entity.
type Address struct {
	ent.Schema
}

// Fields of the Email.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Comment("name | 收件人姓名"),
		field.String("Phone").Comment("phone number | 电话"),
		field.String("Address").Comment("address | 地址详情"),
		field.Uint64("UserID").Optional().Comment("user id | 用户编号"),
	}
}

// Edges of the Email.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_address", User.Type).Ref("address").Field("UserID").Unique(),
	}
}

func (Address) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/schema/mixins"
)

// Email holds the schema definition for the Email entity.
type Email struct {
	ent.Schema
}

// Fields of the Email.
func (Email) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Comment("email | 邮箱号"),
		field.Bool("is_verified").
			Optional().
			Default(false).
			Comment("is verified | 邮箱号是否验证"),
		field.String("secret").Comment("secret | 验证码"),
		field.Uint64("user_id").Comment("user id | user ID"),
	}
}

// Edges of the Email.
func (Email) Edges() []ent.Edge {
	return nil
}

func (Email) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

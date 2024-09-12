package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().Comment("user's login name | 登录名"),
		field.String("password").Comment("password | 密码"),
		field.String("nickname").Unique().Comment("nickname | 昵称"),
		field.String("email").Optional().Comment("email | 邮箱号"),
		field.Uint8("status").Default(1).Optional().Comment("status 1 normal 0 ban | 状态 1 正常 0 禁用"),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("avatar | 头像路径"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

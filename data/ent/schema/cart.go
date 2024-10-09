package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/schema/mixins"
)

type Cart struct {
	ent.Schema
}

//	type Cart struct {
//		gorm.Model
//		UserID    uint
//		ProductID uint
//		Num       uint
//		MaxNum    uint
//		Check     bool
//	}
func (Cart) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("UserID").Optional().Comment("user id | 用户编号"),
		field.Uint64("ProductID").Optional().Comment("product id | 产品编号"),
		field.Int32("Num").Comment("num | 数量"),
		field.Int32("MaxNum").Comment("max num | 最大数量"),
		field.Bool("Check").Comment("check"),
	}
}

func (Cart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_cart", User.Type).Ref("cart").Field("UserID").Unique(),
		edge.From("product_cart", Product.Type).Ref("cart").Field("ProductID").Unique(),
	}
}

func (Cart) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

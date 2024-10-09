package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/schema/mixins"
)

type Order struct {
	ent.Schema
}

//	type Order struct {
//		gorm.Model
//		UserID       uint
//		ProductID    uint
//		Num          uint
//		OrderNum     uint64
//		AddressName  string
//		AddressPhone string
//		Address      string
//		Type         uint
//	}
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("UserID").Optional().Comment("user id | 用户编号"),
		field.Uint64("ProductID").Optional().Comment("product id | 产品编号"),
		field.Int32("Num").Comment("num | 数量"),
		field.Uint64("OrderNum").Comment("order number | 订单号"),
		field.String("AddressName").Comment("name | 收件人姓名"),
		field.String("AddressPhone").Comment("phone number | 电话"),
		field.String("Address").Comment("address | 地址详情"),
		field.Uint64("Type").Comment("Type | 类型"),
		field.Int64("Price").Comment("Price | 价格"),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_order", User.Type).Ref("order").Field("UserID").Unique(),
		edge.From("product_order", Product.Type).Ref("order").Field("ProductID").Unique(),
	}
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

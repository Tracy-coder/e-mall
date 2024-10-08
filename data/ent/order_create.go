// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/order"
	"github.com/Tracy-coder/e-mall/data/ent/product"
	"github.com/Tracy-coder/e-mall/data/ent/user"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrderCreate) SetCreatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrderCreate) SetUpdatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUpdatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetUserID sets the "UserID" field.
func (oc *OrderCreate) SetUserID(u uint64) *OrderCreate {
	oc.mutation.SetUserID(u)
	return oc
}

// SetNillableUserID sets the "UserID" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUserID(u *uint64) *OrderCreate {
	if u != nil {
		oc.SetUserID(*u)
	}
	return oc
}

// SetProductID sets the "ProductID" field.
func (oc *OrderCreate) SetProductID(u uint64) *OrderCreate {
	oc.mutation.SetProductID(u)
	return oc
}

// SetNillableProductID sets the "ProductID" field if the given value is not nil.
func (oc *OrderCreate) SetNillableProductID(u *uint64) *OrderCreate {
	if u != nil {
		oc.SetProductID(*u)
	}
	return oc
}

// SetNum sets the "Num" field.
func (oc *OrderCreate) SetNum(i int32) *OrderCreate {
	oc.mutation.SetNum(i)
	return oc
}

// SetOrderNum sets the "OrderNum" field.
func (oc *OrderCreate) SetOrderNum(u uint64) *OrderCreate {
	oc.mutation.SetOrderNum(u)
	return oc
}

// SetAddressName sets the "AddressName" field.
func (oc *OrderCreate) SetAddressName(s string) *OrderCreate {
	oc.mutation.SetAddressName(s)
	return oc
}

// SetAddressPhone sets the "AddressPhone" field.
func (oc *OrderCreate) SetAddressPhone(s string) *OrderCreate {
	oc.mutation.SetAddressPhone(s)
	return oc
}

// SetAddress sets the "Address" field.
func (oc *OrderCreate) SetAddress(s string) *OrderCreate {
	oc.mutation.SetAddress(s)
	return oc
}

// SetType sets the "Type" field.
func (oc *OrderCreate) SetType(u uint64) *OrderCreate {
	oc.mutation.SetType(u)
	return oc
}

// SetPrice sets the "Price" field.
func (oc *OrderCreate) SetPrice(i int64) *OrderCreate {
	oc.mutation.SetPrice(i)
	return oc
}

// SetID sets the "id" field.
func (oc *OrderCreate) SetID(u uint64) *OrderCreate {
	oc.mutation.SetID(u)
	return oc
}

// SetUserOrderID sets the "user_order" edge to the User entity by ID.
func (oc *OrderCreate) SetUserOrderID(id uint64) *OrderCreate {
	oc.mutation.SetUserOrderID(id)
	return oc
}

// SetNillableUserOrderID sets the "user_order" edge to the User entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillableUserOrderID(id *uint64) *OrderCreate {
	if id != nil {
		oc = oc.SetUserOrderID(*id)
	}
	return oc
}

// SetUserOrder sets the "user_order" edge to the User entity.
func (oc *OrderCreate) SetUserOrder(u *User) *OrderCreate {
	return oc.SetUserOrderID(u.ID)
}

// SetProductOrderID sets the "product_order" edge to the Product entity by ID.
func (oc *OrderCreate) SetProductOrderID(id uint64) *OrderCreate {
	oc.mutation.SetProductOrderID(id)
	return oc
}

// SetNillableProductOrderID sets the "product_order" edge to the Product entity by ID if the given value is not nil.
func (oc *OrderCreate) SetNillableProductOrderID(id *uint64) *OrderCreate {
	if id != nil {
		oc = oc.SetProductOrderID(*id)
	}
	return oc
}

// SetProductOrder sets the "product_order" edge to the Product entity.
func (oc *OrderCreate) SetProductOrder(p *Product) *OrderCreate {
	return oc.SetProductOrderID(p.ID)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := order.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := order.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Order.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Order.updated_at"`)}
	}
	if _, ok := oc.mutation.Num(); !ok {
		return &ValidationError{Name: "Num", err: errors.New(`ent: missing required field "Order.Num"`)}
	}
	if _, ok := oc.mutation.OrderNum(); !ok {
		return &ValidationError{Name: "OrderNum", err: errors.New(`ent: missing required field "Order.OrderNum"`)}
	}
	if _, ok := oc.mutation.AddressName(); !ok {
		return &ValidationError{Name: "AddressName", err: errors.New(`ent: missing required field "Order.AddressName"`)}
	}
	if _, ok := oc.mutation.AddressPhone(); !ok {
		return &ValidationError{Name: "AddressPhone", err: errors.New(`ent: missing required field "Order.AddressPhone"`)}
	}
	if _, ok := oc.mutation.Address(); !ok {
		return &ValidationError{Name: "Address", err: errors.New(`ent: missing required field "Order.Address"`)}
	}
	if _, ok := oc.mutation.GetType(); !ok {
		return &ValidationError{Name: "Type", err: errors.New(`ent: missing required field "Order.Type"`)}
	}
	if _, ok := oc.mutation.Price(); !ok {
		return &ValidationError{Name: "Price", err: errors.New(`ent: missing required field "Order.Price"`)}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(order.Table, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUint64))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oc.mutation.Num(); ok {
		_spec.SetField(order.FieldNum, field.TypeInt32, value)
		_node.Num = value
	}
	if value, ok := oc.mutation.OrderNum(); ok {
		_spec.SetField(order.FieldOrderNum, field.TypeUint64, value)
		_node.OrderNum = value
	}
	if value, ok := oc.mutation.AddressName(); ok {
		_spec.SetField(order.FieldAddressName, field.TypeString, value)
		_node.AddressName = value
	}
	if value, ok := oc.mutation.AddressPhone(); ok {
		_spec.SetField(order.FieldAddressPhone, field.TypeString, value)
		_node.AddressPhone = value
	}
	if value, ok := oc.mutation.Address(); ok {
		_spec.SetField(order.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := oc.mutation.GetType(); ok {
		_spec.SetField(order.FieldType, field.TypeUint64, value)
		_node.Type = value
	}
	if value, ok := oc.mutation.Price(); ok {
		_spec.SetField(order.FieldPrice, field.TypeInt64, value)
		_node.Price = value
	}
	if nodes := oc.mutation.UserOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserOrderTable,
			Columns: []string{order.UserOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.ProductOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.ProductOrderTable,
			Columns: []string{order.ProductOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProductID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	err      error
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}

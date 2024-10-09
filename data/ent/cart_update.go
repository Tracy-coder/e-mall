// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/cart"
	"github.com/Tracy-coder/e-mall/data/ent/predicate"
	"github.com/Tracy-coder/e-mall/data/ent/product"
	"github.com/Tracy-coder/e-mall/data/ent/user"
)

// CartUpdate is the builder for updating Cart entities.
type CartUpdate struct {
	config
	hooks    []Hook
	mutation *CartMutation
}

// Where appends a list predicates to the CartUpdate builder.
func (cu *CartUpdate) Where(ps ...predicate.Cart) *CartUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CartUpdate) SetUpdatedAt(t time.Time) *CartUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetUserID sets the "UserID" field.
func (cu *CartUpdate) SetUserID(u uint64) *CartUpdate {
	cu.mutation.SetUserID(u)
	return cu
}

// SetNillableUserID sets the "UserID" field if the given value is not nil.
func (cu *CartUpdate) SetNillableUserID(u *uint64) *CartUpdate {
	if u != nil {
		cu.SetUserID(*u)
	}
	return cu
}

// ClearUserID clears the value of the "UserID" field.
func (cu *CartUpdate) ClearUserID() *CartUpdate {
	cu.mutation.ClearUserID()
	return cu
}

// SetProductID sets the "ProductID" field.
func (cu *CartUpdate) SetProductID(u uint64) *CartUpdate {
	cu.mutation.SetProductID(u)
	return cu
}

// SetNillableProductID sets the "ProductID" field if the given value is not nil.
func (cu *CartUpdate) SetNillableProductID(u *uint64) *CartUpdate {
	if u != nil {
		cu.SetProductID(*u)
	}
	return cu
}

// ClearProductID clears the value of the "ProductID" field.
func (cu *CartUpdate) ClearProductID() *CartUpdate {
	cu.mutation.ClearProductID()
	return cu
}

// SetNum sets the "Num" field.
func (cu *CartUpdate) SetNum(i int32) *CartUpdate {
	cu.mutation.ResetNum()
	cu.mutation.SetNum(i)
	return cu
}

// SetNillableNum sets the "Num" field if the given value is not nil.
func (cu *CartUpdate) SetNillableNum(i *int32) *CartUpdate {
	if i != nil {
		cu.SetNum(*i)
	}
	return cu
}

// AddNum adds i to the "Num" field.
func (cu *CartUpdate) AddNum(i int32) *CartUpdate {
	cu.mutation.AddNum(i)
	return cu
}

// SetMaxNum sets the "MaxNum" field.
func (cu *CartUpdate) SetMaxNum(i int32) *CartUpdate {
	cu.mutation.ResetMaxNum()
	cu.mutation.SetMaxNum(i)
	return cu
}

// SetNillableMaxNum sets the "MaxNum" field if the given value is not nil.
func (cu *CartUpdate) SetNillableMaxNum(i *int32) *CartUpdate {
	if i != nil {
		cu.SetMaxNum(*i)
	}
	return cu
}

// AddMaxNum adds i to the "MaxNum" field.
func (cu *CartUpdate) AddMaxNum(i int32) *CartUpdate {
	cu.mutation.AddMaxNum(i)
	return cu
}

// SetCheck sets the "Check" field.
func (cu *CartUpdate) SetCheck(b bool) *CartUpdate {
	cu.mutation.SetCheck(b)
	return cu
}

// SetNillableCheck sets the "Check" field if the given value is not nil.
func (cu *CartUpdate) SetNillableCheck(b *bool) *CartUpdate {
	if b != nil {
		cu.SetCheck(*b)
	}
	return cu
}

// SetUserCartID sets the "user_cart" edge to the User entity by ID.
func (cu *CartUpdate) SetUserCartID(id uint64) *CartUpdate {
	cu.mutation.SetUserCartID(id)
	return cu
}

// SetNillableUserCartID sets the "user_cart" edge to the User entity by ID if the given value is not nil.
func (cu *CartUpdate) SetNillableUserCartID(id *uint64) *CartUpdate {
	if id != nil {
		cu = cu.SetUserCartID(*id)
	}
	return cu
}

// SetUserCart sets the "user_cart" edge to the User entity.
func (cu *CartUpdate) SetUserCart(u *User) *CartUpdate {
	return cu.SetUserCartID(u.ID)
}

// SetProductCartID sets the "product_cart" edge to the Product entity by ID.
func (cu *CartUpdate) SetProductCartID(id uint64) *CartUpdate {
	cu.mutation.SetProductCartID(id)
	return cu
}

// SetNillableProductCartID sets the "product_cart" edge to the Product entity by ID if the given value is not nil.
func (cu *CartUpdate) SetNillableProductCartID(id *uint64) *CartUpdate {
	if id != nil {
		cu = cu.SetProductCartID(*id)
	}
	return cu
}

// SetProductCart sets the "product_cart" edge to the Product entity.
func (cu *CartUpdate) SetProductCart(p *Product) *CartUpdate {
	return cu.SetProductCartID(p.ID)
}

// Mutation returns the CartMutation object of the builder.
func (cu *CartUpdate) Mutation() *CartMutation {
	return cu.mutation
}

// ClearUserCart clears the "user_cart" edge to the User entity.
func (cu *CartUpdate) ClearUserCart() *CartUpdate {
	cu.mutation.ClearUserCart()
	return cu
}

// ClearProductCart clears the "product_cart" edge to the Product entity.
func (cu *CartUpdate) ClearProductCart() *CartUpdate {
	cu.mutation.ClearProductCart()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CartUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CartUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CartUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CartUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CartUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := cart.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CartUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(cart.Table, cart.Columns, sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUint64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(cart.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Num(); ok {
		_spec.SetField(cart.FieldNum, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.AddedNum(); ok {
		_spec.AddField(cart.FieldNum, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.MaxNum(); ok {
		_spec.SetField(cart.FieldMaxNum, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.AddedMaxNum(); ok {
		_spec.AddField(cart.FieldMaxNum, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.Check(); ok {
		_spec.SetField(cart.FieldCheck, field.TypeBool, value)
	}
	if cu.mutation.UserCartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UserCartTable,
			Columns: []string{cart.UserCartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserCartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UserCartTable,
			Columns: []string{cart.UserCartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ProductCartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.ProductCartTable,
			Columns: []string{cart.ProductCartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ProductCartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.ProductCartTable,
			Columns: []string{cart.ProductCartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CartUpdateOne is the builder for updating a single Cart entity.
type CartUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CartMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CartUpdateOne) SetUpdatedAt(t time.Time) *CartUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetUserID sets the "UserID" field.
func (cuo *CartUpdateOne) SetUserID(u uint64) *CartUpdateOne {
	cuo.mutation.SetUserID(u)
	return cuo
}

// SetNillableUserID sets the "UserID" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableUserID(u *uint64) *CartUpdateOne {
	if u != nil {
		cuo.SetUserID(*u)
	}
	return cuo
}

// ClearUserID clears the value of the "UserID" field.
func (cuo *CartUpdateOne) ClearUserID() *CartUpdateOne {
	cuo.mutation.ClearUserID()
	return cuo
}

// SetProductID sets the "ProductID" field.
func (cuo *CartUpdateOne) SetProductID(u uint64) *CartUpdateOne {
	cuo.mutation.SetProductID(u)
	return cuo
}

// SetNillableProductID sets the "ProductID" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableProductID(u *uint64) *CartUpdateOne {
	if u != nil {
		cuo.SetProductID(*u)
	}
	return cuo
}

// ClearProductID clears the value of the "ProductID" field.
func (cuo *CartUpdateOne) ClearProductID() *CartUpdateOne {
	cuo.mutation.ClearProductID()
	return cuo
}

// SetNum sets the "Num" field.
func (cuo *CartUpdateOne) SetNum(i int32) *CartUpdateOne {
	cuo.mutation.ResetNum()
	cuo.mutation.SetNum(i)
	return cuo
}

// SetNillableNum sets the "Num" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableNum(i *int32) *CartUpdateOne {
	if i != nil {
		cuo.SetNum(*i)
	}
	return cuo
}

// AddNum adds i to the "Num" field.
func (cuo *CartUpdateOne) AddNum(i int32) *CartUpdateOne {
	cuo.mutation.AddNum(i)
	return cuo
}

// SetMaxNum sets the "MaxNum" field.
func (cuo *CartUpdateOne) SetMaxNum(i int32) *CartUpdateOne {
	cuo.mutation.ResetMaxNum()
	cuo.mutation.SetMaxNum(i)
	return cuo
}

// SetNillableMaxNum sets the "MaxNum" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableMaxNum(i *int32) *CartUpdateOne {
	if i != nil {
		cuo.SetMaxNum(*i)
	}
	return cuo
}

// AddMaxNum adds i to the "MaxNum" field.
func (cuo *CartUpdateOne) AddMaxNum(i int32) *CartUpdateOne {
	cuo.mutation.AddMaxNum(i)
	return cuo
}

// SetCheck sets the "Check" field.
func (cuo *CartUpdateOne) SetCheck(b bool) *CartUpdateOne {
	cuo.mutation.SetCheck(b)
	return cuo
}

// SetNillableCheck sets the "Check" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableCheck(b *bool) *CartUpdateOne {
	if b != nil {
		cuo.SetCheck(*b)
	}
	return cuo
}

// SetUserCartID sets the "user_cart" edge to the User entity by ID.
func (cuo *CartUpdateOne) SetUserCartID(id uint64) *CartUpdateOne {
	cuo.mutation.SetUserCartID(id)
	return cuo
}

// SetNillableUserCartID sets the "user_cart" edge to the User entity by ID if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableUserCartID(id *uint64) *CartUpdateOne {
	if id != nil {
		cuo = cuo.SetUserCartID(*id)
	}
	return cuo
}

// SetUserCart sets the "user_cart" edge to the User entity.
func (cuo *CartUpdateOne) SetUserCart(u *User) *CartUpdateOne {
	return cuo.SetUserCartID(u.ID)
}

// SetProductCartID sets the "product_cart" edge to the Product entity by ID.
func (cuo *CartUpdateOne) SetProductCartID(id uint64) *CartUpdateOne {
	cuo.mutation.SetProductCartID(id)
	return cuo
}

// SetNillableProductCartID sets the "product_cart" edge to the Product entity by ID if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableProductCartID(id *uint64) *CartUpdateOne {
	if id != nil {
		cuo = cuo.SetProductCartID(*id)
	}
	return cuo
}

// SetProductCart sets the "product_cart" edge to the Product entity.
func (cuo *CartUpdateOne) SetProductCart(p *Product) *CartUpdateOne {
	return cuo.SetProductCartID(p.ID)
}

// Mutation returns the CartMutation object of the builder.
func (cuo *CartUpdateOne) Mutation() *CartMutation {
	return cuo.mutation
}

// ClearUserCart clears the "user_cart" edge to the User entity.
func (cuo *CartUpdateOne) ClearUserCart() *CartUpdateOne {
	cuo.mutation.ClearUserCart()
	return cuo
}

// ClearProductCart clears the "product_cart" edge to the Product entity.
func (cuo *CartUpdateOne) ClearProductCart() *CartUpdateOne {
	cuo.mutation.ClearProductCart()
	return cuo
}

// Where appends a list predicates to the CartUpdate builder.
func (cuo *CartUpdateOne) Where(ps ...predicate.Cart) *CartUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CartUpdateOne) Select(field string, fields ...string) *CartUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cart entity.
func (cuo *CartUpdateOne) Save(ctx context.Context) (*Cart, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CartUpdateOne) SaveX(ctx context.Context) *Cart {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CartUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CartUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CartUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := cart.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CartUpdateOne) sqlSave(ctx context.Context) (_node *Cart, err error) {
	_spec := sqlgraph.NewUpdateSpec(cart.Table, cart.Columns, sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUint64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cart.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cart.FieldID)
		for _, f := range fields {
			if !cart.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cart.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cart.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Num(); ok {
		_spec.SetField(cart.FieldNum, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.AddedNum(); ok {
		_spec.AddField(cart.FieldNum, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.MaxNum(); ok {
		_spec.SetField(cart.FieldMaxNum, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.AddedMaxNum(); ok {
		_spec.AddField(cart.FieldMaxNum, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.Check(); ok {
		_spec.SetField(cart.FieldCheck, field.TypeBool, value)
	}
	if cuo.mutation.UserCartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UserCartTable,
			Columns: []string{cart.UserCartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserCartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UserCartTable,
			Columns: []string{cart.UserCartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ProductCartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.ProductCartTable,
			Columns: []string{cart.ProductCartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ProductCartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.ProductCartTable,
			Columns: []string{cart.ProductCartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Cart{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}

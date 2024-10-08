// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/carousel"
	"github.com/Tracy-coder/e-mall/data/ent/product"
)

// CarouselCreate is the builder for creating a Carousel entity.
type CarouselCreate struct {
	config
	mutation *CarouselMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CarouselCreate) SetCreatedAt(t time.Time) *CarouselCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CarouselCreate) SetNillableCreatedAt(t *time.Time) *CarouselCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CarouselCreate) SetUpdatedAt(t time.Time) *CarouselCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CarouselCreate) SetNillableUpdatedAt(t *time.Time) *CarouselCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetProductID sets the "ProductID" field.
func (cc *CarouselCreate) SetProductID(u uint64) *CarouselCreate {
	cc.mutation.SetProductID(u)
	return cc
}

// SetNillableProductID sets the "ProductID" field if the given value is not nil.
func (cc *CarouselCreate) SetNillableProductID(u *uint64) *CarouselCreate {
	if u != nil {
		cc.SetProductID(*u)
	}
	return cc
}

// SetImgPath sets the "ImgPath" field.
func (cc *CarouselCreate) SetImgPath(s string) *CarouselCreate {
	cc.mutation.SetImgPath(s)
	return cc
}

// SetID sets the "id" field.
func (cc *CarouselCreate) SetID(u uint64) *CarouselCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetOwnerID sets the "owner" edge to the Product entity by ID.
func (cc *CarouselCreate) SetOwnerID(id uint64) *CarouselCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetNillableOwnerID sets the "owner" edge to the Product entity by ID if the given value is not nil.
func (cc *CarouselCreate) SetNillableOwnerID(id *uint64) *CarouselCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the "owner" edge to the Product entity.
func (cc *CarouselCreate) SetOwner(p *Product) *CarouselCreate {
	return cc.SetOwnerID(p.ID)
}

// Mutation returns the CarouselMutation object of the builder.
func (cc *CarouselCreate) Mutation() *CarouselMutation {
	return cc.mutation
}

// Save creates the Carousel in the database.
func (cc *CarouselCreate) Save(ctx context.Context) (*Carousel, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CarouselCreate) SaveX(ctx context.Context) *Carousel {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CarouselCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CarouselCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CarouselCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := carousel.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := carousel.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CarouselCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Carousel.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Carousel.updated_at"`)}
	}
	if _, ok := cc.mutation.ImgPath(); !ok {
		return &ValidationError{Name: "ImgPath", err: errors.New(`ent: missing required field "Carousel.ImgPath"`)}
	}
	return nil
}

func (cc *CarouselCreate) sqlSave(ctx context.Context) (*Carousel, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CarouselCreate) createSpec() (*Carousel, *sqlgraph.CreateSpec) {
	var (
		_node = &Carousel{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(carousel.Table, sqlgraph.NewFieldSpec(carousel.FieldID, field.TypeUint64))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(carousel.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(carousel.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.ImgPath(); ok {
		_spec.SetField(carousel.FieldImgPath, field.TypeString, value)
		_node.ImgPath = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   carousel.OwnerTable,
			Columns: []string{carousel.OwnerColumn},
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

// CarouselCreateBulk is the builder for creating many Carousel entities in bulk.
type CarouselCreateBulk struct {
	config
	err      error
	builders []*CarouselCreate
}

// Save creates the Carousel entities in the database.
func (ccb *CarouselCreateBulk) Save(ctx context.Context) ([]*Carousel, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Carousel, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CarouselMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CarouselCreateBulk) SaveX(ctx context.Context) []*Carousel {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CarouselCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CarouselCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/mysql/ent/link"
)

// LinkCreate is the builder for creating a Link entity.
type LinkCreate struct {
	config
	mutation *LinkMutation
	hooks    []Hook
}

// SetURL sets the "url" field.
func (lc *LinkCreate) SetURL(s string) *LinkCreate {
	lc.mutation.SetURL(s)
	return lc
}

// SetHash sets the "hash" field.
func (lc *LinkCreate) SetHash(s string) *LinkCreate {
	lc.mutation.SetHash(s)
	return lc
}

// SetDescribe sets the "describe" field.
func (lc *LinkCreate) SetDescribe(s string) *LinkCreate {
	lc.mutation.SetDescribe(s)
	return lc
}

// SetNillableDescribe sets the "describe" field if the given value is not nil.
func (lc *LinkCreate) SetNillableDescribe(s *string) *LinkCreate {
	if s != nil {
		lc.SetDescribe(*s)
	}
	return lc
}

// SetJSON sets the "json" field.
func (lc *LinkCreate) SetJSON(v v1.Link) *LinkCreate {
	lc.mutation.SetJSON(v)
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *LinkCreate) SetCreatedAt(t time.Time) *LinkCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LinkCreate) SetUpdatedAt(t time.Time) *LinkCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetID sets the "id" field.
func (lc *LinkCreate) SetID(u uuid.UUID) *LinkCreate {
	lc.mutation.SetID(u)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LinkCreate) SetNillableID(u *uuid.UUID) *LinkCreate {
	if u != nil {
		lc.SetID(*u)
	}
	return lc
}

// Mutation returns the LinkMutation object of the builder.
func (lc *LinkCreate) Mutation() *LinkMutation {
	return lc.mutation
}

// Save creates the Link in the database.
func (lc *LinkCreate) Save(ctx context.Context) (*Link, error) {
	lc.defaults()
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LinkCreate) SaveX(ctx context.Context) *Link {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LinkCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LinkCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LinkCreate) defaults() {
	if _, ok := lc.mutation.ID(); !ok {
		v := link.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LinkCreate) check() error {
	if _, ok := lc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Link.url"`)}
	}
	if v, ok := lc.mutation.URL(); ok {
		if err := link.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Link.url": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Link.hash"`)}
	}
	if v, ok := lc.mutation.Hash(); ok {
		if err := link.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "Link.hash": %w`, err)}
		}
	}
	if _, ok := lc.mutation.JSON(); !ok {
		return &ValidationError{Name: "json", err: errors.New(`ent: missing required field "Link.json"`)}
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Link.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Link.updated_at"`)}
	}
	return nil
}

func (lc *LinkCreate) sqlSave(ctx context.Context) (*Link, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LinkCreate) createSpec() (*Link, *sqlgraph.CreateSpec) {
	var (
		_node = &Link{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(link.Table, sqlgraph.NewFieldSpec(link.FieldID, field.TypeUUID))
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lc.mutation.URL(); ok {
		_spec.SetField(link.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := lc.mutation.Hash(); ok {
		_spec.SetField(link.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if value, ok := lc.mutation.Describe(); ok {
		_spec.SetField(link.FieldDescribe, field.TypeString, value)
		_node.Describe = value
	}
	if value, ok := lc.mutation.JSON(); ok {
		_spec.SetField(link.FieldJSON, field.TypeJSON, value)
		_node.JSON = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(link.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(link.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// LinkCreateBulk is the builder for creating many Link entities in bulk.
type LinkCreateBulk struct {
	config
	err      error
	builders []*LinkCreate
}

// Save creates the Link entities in the database.
func (lcb *LinkCreateBulk) Save(ctx context.Context) ([]*Link, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Link, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinkMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LinkCreateBulk) SaveX(ctx context.Context) []*Link {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LinkCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LinkCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
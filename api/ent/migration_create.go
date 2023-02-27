// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuoitrevohoc/app-template/api/ent/migration"
)

// MigrationCreate is the builder for creating a Migration entity.
type MigrationCreate struct {
	config
	mutation *MigrationMutation
	hooks    []Hook
}

// SetMigration sets the "migration" field.
func (mc *MigrationCreate) SetMigration(s string) *MigrationCreate {
	mc.mutation.SetMigration(s)
	return mc
}

// SetExecutionAt sets the "execution_at" field.
func (mc *MigrationCreate) SetExecutionAt(t time.Time) *MigrationCreate {
	mc.mutation.SetExecutionAt(t)
	return mc
}

// SetNillableExecutionAt sets the "execution_at" field if the given value is not nil.
func (mc *MigrationCreate) SetNillableExecutionAt(t *time.Time) *MigrationCreate {
	if t != nil {
		mc.SetExecutionAt(*t)
	}
	return mc
}

// Mutation returns the MigrationMutation object of the builder.
func (mc *MigrationCreate) Mutation() *MigrationMutation {
	return mc.mutation
}

// Save creates the Migration in the database.
func (mc *MigrationCreate) Save(ctx context.Context) (*Migration, error) {
	mc.defaults()
	return withHooks[*Migration, MigrationMutation](ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MigrationCreate) SaveX(ctx context.Context) *Migration {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MigrationCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MigrationCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MigrationCreate) defaults() {
	if _, ok := mc.mutation.ExecutionAt(); !ok {
		v := migration.DefaultExecutionAt()
		mc.mutation.SetExecutionAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MigrationCreate) check() error {
	if _, ok := mc.mutation.Migration(); !ok {
		return &ValidationError{Name: "migration", err: errors.New(`ent: missing required field "Migration.migration"`)}
	}
	if _, ok := mc.mutation.ExecutionAt(); !ok {
		return &ValidationError{Name: "execution_at", err: errors.New(`ent: missing required field "Migration.execution_at"`)}
	}
	return nil
}

func (mc *MigrationCreate) sqlSave(ctx context.Context) (*Migration, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MigrationCreate) createSpec() (*Migration, *sqlgraph.CreateSpec) {
	var (
		_node = &Migration{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: migration.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: migration.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Migration(); ok {
		_spec.SetField(migration.FieldMigration, field.TypeString, value)
		_node.Migration = value
	}
	if value, ok := mc.mutation.ExecutionAt(); ok {
		_spec.SetField(migration.FieldExecutionAt, field.TypeTime, value)
		_node.ExecutionAt = value
	}
	return _node, _spec
}

// MigrationCreateBulk is the builder for creating many Migration entities in bulk.
type MigrationCreateBulk struct {
	config
	builders []*MigrationCreate
}

// Save creates the Migration entities in the database.
func (mcb *MigrationCreateBulk) Save(ctx context.Context) ([]*Migration, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Migration, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MigrationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MigrationCreateBulk) SaveX(ctx context.Context) []*Migration {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MigrationCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MigrationCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuoitrevohoc/app-template/backend/pkg/ent/invoice"
)

// InvoiceCreate is the builder for creating a Invoice entity.
type InvoiceCreate struct {
	config
	mutation *InvoiceMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (ic *InvoiceCreate) SetTitle(s string) *InvoiceCreate {
	ic.mutation.SetTitle(s)
	return ic
}

// SetLeetCodeLink sets the "leet_code_link" field.
func (ic *InvoiceCreate) SetLeetCodeLink(s string) *InvoiceCreate {
	ic.mutation.SetLeetCodeLink(s)
	return ic
}

// SetInvoicedTo sets the "invoiced_to" field.
func (ic *InvoiceCreate) SetInvoicedTo(s string) *InvoiceCreate {
	ic.mutation.SetInvoicedTo(s)
	return ic
}

// Mutation returns the InvoiceMutation object of the builder.
func (ic *InvoiceCreate) Mutation() *InvoiceMutation {
	return ic.mutation
}

// Save creates the Invoice in the database.
func (ic *InvoiceCreate) Save(ctx context.Context) (*Invoice, error) {
	return withHooks[*Invoice, InvoiceMutation](ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InvoiceCreate) SaveX(ctx context.Context) *Invoice {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InvoiceCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InvoiceCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InvoiceCreate) check() error {
	if _, ok := ic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Invoice.title"`)}
	}
	if _, ok := ic.mutation.LeetCodeLink(); !ok {
		return &ValidationError{Name: "leet_code_link", err: errors.New(`ent: missing required field "Invoice.leet_code_link"`)}
	}
	if _, ok := ic.mutation.InvoicedTo(); !ok {
		return &ValidationError{Name: "invoiced_to", err: errors.New(`ent: missing required field "Invoice.invoiced_to"`)}
	}
	return nil
}

func (ic *InvoiceCreate) sqlSave(ctx context.Context) (*Invoice, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *InvoiceCreate) createSpec() (*Invoice, *sqlgraph.CreateSpec) {
	var (
		_node = &Invoice{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: invoice.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invoice.FieldID,
			},
		}
	)
	if value, ok := ic.mutation.Title(); ok {
		_spec.SetField(invoice.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ic.mutation.LeetCodeLink(); ok {
		_spec.SetField(invoice.FieldLeetCodeLink, field.TypeString, value)
		_node.LeetCodeLink = value
	}
	if value, ok := ic.mutation.InvoicedTo(); ok {
		_spec.SetField(invoice.FieldInvoicedTo, field.TypeString, value)
		_node.InvoicedTo = value
	}
	return _node, _spec
}

// InvoiceCreateBulk is the builder for creating many Invoice entities in bulk.
type InvoiceCreateBulk struct {
	config
	builders []*InvoiceCreate
}

// Save creates the Invoice entities in the database.
func (icb *InvoiceCreateBulk) Save(ctx context.Context) ([]*Invoice, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Invoice, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvoiceMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InvoiceCreateBulk) SaveX(ctx context.Context) []*Invoice {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InvoiceCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InvoiceCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/tuoitrevohoc/app-template/api/ent/invoice"
	"github.com/tuoitrevohoc/app-template/api/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeInvoice = "Invoice"
)

// InvoiceMutation represents an operation that mutates the Invoice nodes in the graph.
type InvoiceMutation struct {
	config
	op             Op
	typ            string
	id             *int
	title          *string
	leet_code_link *string
	invoiced_to    *string
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Invoice, error)
	predicates     []predicate.Invoice
}

var _ ent.Mutation = (*InvoiceMutation)(nil)

// invoiceOption allows management of the mutation configuration using functional options.
type invoiceOption func(*InvoiceMutation)

// newInvoiceMutation creates new mutation for the Invoice entity.
func newInvoiceMutation(c config, op Op, opts ...invoiceOption) *InvoiceMutation {
	m := &InvoiceMutation{
		config:        c,
		op:            op,
		typ:           TypeInvoice,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withInvoiceID sets the ID field of the mutation.
func withInvoiceID(id int) invoiceOption {
	return func(m *InvoiceMutation) {
		var (
			err   error
			once  sync.Once
			value *Invoice
		)
		m.oldValue = func(ctx context.Context) (*Invoice, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Invoice.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withInvoice sets the old Invoice of the mutation.
func withInvoice(node *Invoice) invoiceOption {
	return func(m *InvoiceMutation) {
		m.oldValue = func(context.Context) (*Invoice, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m InvoiceMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m InvoiceMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *InvoiceMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *InvoiceMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Invoice.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *InvoiceMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *InvoiceMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Invoice entity.
// If the Invoice object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *InvoiceMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *InvoiceMutation) ResetTitle() {
	m.title = nil
}

// SetLeetCodeLink sets the "leet_code_link" field.
func (m *InvoiceMutation) SetLeetCodeLink(s string) {
	m.leet_code_link = &s
}

// LeetCodeLink returns the value of the "leet_code_link" field in the mutation.
func (m *InvoiceMutation) LeetCodeLink() (r string, exists bool) {
	v := m.leet_code_link
	if v == nil {
		return
	}
	return *v, true
}

// OldLeetCodeLink returns the old "leet_code_link" field's value of the Invoice entity.
// If the Invoice object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *InvoiceMutation) OldLeetCodeLink(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLeetCodeLink is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLeetCodeLink requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLeetCodeLink: %w", err)
	}
	return oldValue.LeetCodeLink, nil
}

// ResetLeetCodeLink resets all changes to the "leet_code_link" field.
func (m *InvoiceMutation) ResetLeetCodeLink() {
	m.leet_code_link = nil
}

// SetInvoicedTo sets the "invoiced_to" field.
func (m *InvoiceMutation) SetInvoicedTo(s string) {
	m.invoiced_to = &s
}

// InvoicedTo returns the value of the "invoiced_to" field in the mutation.
func (m *InvoiceMutation) InvoicedTo() (r string, exists bool) {
	v := m.invoiced_to
	if v == nil {
		return
	}
	return *v, true
}

// OldInvoicedTo returns the old "invoiced_to" field's value of the Invoice entity.
// If the Invoice object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *InvoiceMutation) OldInvoicedTo(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldInvoicedTo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldInvoicedTo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInvoicedTo: %w", err)
	}
	return oldValue.InvoicedTo, nil
}

// ResetInvoicedTo resets all changes to the "invoiced_to" field.
func (m *InvoiceMutation) ResetInvoicedTo() {
	m.invoiced_to = nil
}

// Where appends a list predicates to the InvoiceMutation builder.
func (m *InvoiceMutation) Where(ps ...predicate.Invoice) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the InvoiceMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *InvoiceMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Invoice, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *InvoiceMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *InvoiceMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Invoice).
func (m *InvoiceMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *InvoiceMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.title != nil {
		fields = append(fields, invoice.FieldTitle)
	}
	if m.leet_code_link != nil {
		fields = append(fields, invoice.FieldLeetCodeLink)
	}
	if m.invoiced_to != nil {
		fields = append(fields, invoice.FieldInvoicedTo)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *InvoiceMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case invoice.FieldTitle:
		return m.Title()
	case invoice.FieldLeetCodeLink:
		return m.LeetCodeLink()
	case invoice.FieldInvoicedTo:
		return m.InvoicedTo()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *InvoiceMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case invoice.FieldTitle:
		return m.OldTitle(ctx)
	case invoice.FieldLeetCodeLink:
		return m.OldLeetCodeLink(ctx)
	case invoice.FieldInvoicedTo:
		return m.OldInvoicedTo(ctx)
	}
	return nil, fmt.Errorf("unknown Invoice field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *InvoiceMutation) SetField(name string, value ent.Value) error {
	switch name {
	case invoice.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case invoice.FieldLeetCodeLink:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLeetCodeLink(v)
		return nil
	case invoice.FieldInvoicedTo:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInvoicedTo(v)
		return nil
	}
	return fmt.Errorf("unknown Invoice field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *InvoiceMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *InvoiceMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *InvoiceMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Invoice numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *InvoiceMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *InvoiceMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *InvoiceMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Invoice nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *InvoiceMutation) ResetField(name string) error {
	switch name {
	case invoice.FieldTitle:
		m.ResetTitle()
		return nil
	case invoice.FieldLeetCodeLink:
		m.ResetLeetCodeLink()
		return nil
	case invoice.FieldInvoicedTo:
		m.ResetInvoicedTo()
		return nil
	}
	return fmt.Errorf("unknown Invoice field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *InvoiceMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *InvoiceMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *InvoiceMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *InvoiceMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *InvoiceMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *InvoiceMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *InvoiceMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Invoice unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *InvoiceMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Invoice edge %s", name)
}

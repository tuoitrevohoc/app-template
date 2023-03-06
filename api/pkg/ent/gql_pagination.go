// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/tuoitrevohoc/app-template/api/pkg/ent/invoice"
	"github.com/tuoitrevohoc/app-template/api/pkg/ent/migration"
	"github.com/tuoitrevohoc/app-template/api/pkg/ent/permission"
	"github.com/tuoitrevohoc/app-template/api/pkg/ent/role"
	"github.com/tuoitrevohoc/app-template/api/pkg/ent/user"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// InvoiceEdge is the edge representation of Invoice.
type InvoiceEdge struct {
	Node   *Invoice `json:"node"`
	Cursor Cursor   `json:"cursor"`
}

// InvoiceConnection is the connection containing edges to Invoice.
type InvoiceConnection struct {
	Edges      []*InvoiceEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

func (c *InvoiceConnection) build(nodes []*Invoice, pager *invoicePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Invoice
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Invoice {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Invoice {
			return nodes[i]
		}
	}
	c.Edges = make([]*InvoiceEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &InvoiceEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// InvoicePaginateOption enables pagination customization.
type InvoicePaginateOption func(*invoicePager) error

// WithInvoiceOrder configures pagination ordering.
func WithInvoiceOrder(order *InvoiceOrder) InvoicePaginateOption {
	if order == nil {
		order = DefaultInvoiceOrder
	}
	o := *order
	return func(pager *invoicePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultInvoiceOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithInvoiceFilter configures pagination filter.
func WithInvoiceFilter(filter func(*InvoiceQuery) (*InvoiceQuery, error)) InvoicePaginateOption {
	return func(pager *invoicePager) error {
		if filter == nil {
			return errors.New("InvoiceQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type invoicePager struct {
	order  *InvoiceOrder
	filter func(*InvoiceQuery) (*InvoiceQuery, error)
}

func newInvoicePager(opts []InvoicePaginateOption) (*invoicePager, error) {
	pager := &invoicePager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultInvoiceOrder
	}
	return pager, nil
}

func (p *invoicePager) applyFilter(query *InvoiceQuery) (*InvoiceQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *invoicePager) toCursor(i *Invoice) Cursor {
	return p.order.Field.toCursor(i)
}

func (p *invoicePager) applyCursors(query *InvoiceQuery, after, before *Cursor) *InvoiceQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultInvoiceOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *invoicePager) applyOrder(query *InvoiceQuery, reverse bool) *InvoiceQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultInvoiceOrder.Field {
		query = query.Order(direction.orderFunc(DefaultInvoiceOrder.Field.field))
	}
	return query
}

func (p *invoicePager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultInvoiceOrder.Field {
			b.Comma().Ident(DefaultInvoiceOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Invoice.
func (i *InvoiceQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...InvoicePaginateOption,
) (*InvoiceConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newInvoicePager(opts)
	if err != nil {
		return nil, err
	}
	if i, err = pager.applyFilter(i); err != nil {
		return nil, err
	}
	conn := &InvoiceConnection{Edges: []*InvoiceEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = i.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	i = pager.applyCursors(i, after, before)
	i = pager.applyOrder(i, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		i.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := i.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := i.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// InvoiceOrderField defines the ordering field of Invoice.
type InvoiceOrderField struct {
	field    string
	toCursor func(*Invoice) Cursor
}

// InvoiceOrder defines the ordering of Invoice.
type InvoiceOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     *InvoiceOrderField `json:"field"`
}

// DefaultInvoiceOrder is the default ordering of Invoice.
var DefaultInvoiceOrder = &InvoiceOrder{
	Direction: OrderDirectionAsc,
	Field: &InvoiceOrderField{
		field: invoice.FieldID,
		toCursor: func(i *Invoice) Cursor {
			return Cursor{ID: i.ID}
		},
	},
}

// ToEdge converts Invoice into InvoiceEdge.
func (i *Invoice) ToEdge(order *InvoiceOrder) *InvoiceEdge {
	if order == nil {
		order = DefaultInvoiceOrder
	}
	return &InvoiceEdge{
		Node:   i,
		Cursor: order.Field.toCursor(i),
	}
}

// MigrationEdge is the edge representation of Migration.
type MigrationEdge struct {
	Node   *Migration `json:"node"`
	Cursor Cursor     `json:"cursor"`
}

// MigrationConnection is the connection containing edges to Migration.
type MigrationConnection struct {
	Edges      []*MigrationEdge `json:"edges"`
	PageInfo   PageInfo         `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

func (c *MigrationConnection) build(nodes []*Migration, pager *migrationPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Migration
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Migration {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Migration {
			return nodes[i]
		}
	}
	c.Edges = make([]*MigrationEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &MigrationEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// MigrationPaginateOption enables pagination customization.
type MigrationPaginateOption func(*migrationPager) error

// WithMigrationOrder configures pagination ordering.
func WithMigrationOrder(order *MigrationOrder) MigrationPaginateOption {
	if order == nil {
		order = DefaultMigrationOrder
	}
	o := *order
	return func(pager *migrationPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultMigrationOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithMigrationFilter configures pagination filter.
func WithMigrationFilter(filter func(*MigrationQuery) (*MigrationQuery, error)) MigrationPaginateOption {
	return func(pager *migrationPager) error {
		if filter == nil {
			return errors.New("MigrationQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type migrationPager struct {
	order  *MigrationOrder
	filter func(*MigrationQuery) (*MigrationQuery, error)
}

func newMigrationPager(opts []MigrationPaginateOption) (*migrationPager, error) {
	pager := &migrationPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultMigrationOrder
	}
	return pager, nil
}

func (p *migrationPager) applyFilter(query *MigrationQuery) (*MigrationQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *migrationPager) toCursor(m *Migration) Cursor {
	return p.order.Field.toCursor(m)
}

func (p *migrationPager) applyCursors(query *MigrationQuery, after, before *Cursor) *MigrationQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultMigrationOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *migrationPager) applyOrder(query *MigrationQuery, reverse bool) *MigrationQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultMigrationOrder.Field {
		query = query.Order(direction.orderFunc(DefaultMigrationOrder.Field.field))
	}
	return query
}

func (p *migrationPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultMigrationOrder.Field {
			b.Comma().Ident(DefaultMigrationOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Migration.
func (m *MigrationQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...MigrationPaginateOption,
) (*MigrationConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newMigrationPager(opts)
	if err != nil {
		return nil, err
	}
	if m, err = pager.applyFilter(m); err != nil {
		return nil, err
	}
	conn := &MigrationConnection{Edges: []*MigrationEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = m.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	m = pager.applyCursors(m, after, before)
	m = pager.applyOrder(m, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		m.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := m.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := m.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// MigrationOrderField defines the ordering field of Migration.
type MigrationOrderField struct {
	field    string
	toCursor func(*Migration) Cursor
}

// MigrationOrder defines the ordering of Migration.
type MigrationOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     *MigrationOrderField `json:"field"`
}

// DefaultMigrationOrder is the default ordering of Migration.
var DefaultMigrationOrder = &MigrationOrder{
	Direction: OrderDirectionAsc,
	Field: &MigrationOrderField{
		field: migration.FieldID,
		toCursor: func(m *Migration) Cursor {
			return Cursor{ID: m.ID}
		},
	},
}

// ToEdge converts Migration into MigrationEdge.
func (m *Migration) ToEdge(order *MigrationOrder) *MigrationEdge {
	if order == nil {
		order = DefaultMigrationOrder
	}
	return &MigrationEdge{
		Node:   m,
		Cursor: order.Field.toCursor(m),
	}
}

// PermissionEdge is the edge representation of Permission.
type PermissionEdge struct {
	Node   *Permission `json:"node"`
	Cursor Cursor      `json:"cursor"`
}

// PermissionConnection is the connection containing edges to Permission.
type PermissionConnection struct {
	Edges      []*PermissionEdge `json:"edges"`
	PageInfo   PageInfo          `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

func (c *PermissionConnection) build(nodes []*Permission, pager *permissionPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Permission
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Permission {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Permission {
			return nodes[i]
		}
	}
	c.Edges = make([]*PermissionEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &PermissionEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// PermissionPaginateOption enables pagination customization.
type PermissionPaginateOption func(*permissionPager) error

// WithPermissionOrder configures pagination ordering.
func WithPermissionOrder(order *PermissionOrder) PermissionPaginateOption {
	if order == nil {
		order = DefaultPermissionOrder
	}
	o := *order
	return func(pager *permissionPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultPermissionOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithPermissionFilter configures pagination filter.
func WithPermissionFilter(filter func(*PermissionQuery) (*PermissionQuery, error)) PermissionPaginateOption {
	return func(pager *permissionPager) error {
		if filter == nil {
			return errors.New("PermissionQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type permissionPager struct {
	order  *PermissionOrder
	filter func(*PermissionQuery) (*PermissionQuery, error)
}

func newPermissionPager(opts []PermissionPaginateOption) (*permissionPager, error) {
	pager := &permissionPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultPermissionOrder
	}
	return pager, nil
}

func (p *permissionPager) applyFilter(query *PermissionQuery) (*PermissionQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *permissionPager) toCursor(pe *Permission) Cursor {
	return p.order.Field.toCursor(pe)
}

func (p *permissionPager) applyCursors(query *PermissionQuery, after, before *Cursor) *PermissionQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultPermissionOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *permissionPager) applyOrder(query *PermissionQuery, reverse bool) *PermissionQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultPermissionOrder.Field {
		query = query.Order(direction.orderFunc(DefaultPermissionOrder.Field.field))
	}
	return query
}

func (p *permissionPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultPermissionOrder.Field {
			b.Comma().Ident(DefaultPermissionOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Permission.
func (pe *PermissionQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...PermissionPaginateOption,
) (*PermissionConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newPermissionPager(opts)
	if err != nil {
		return nil, err
	}
	if pe, err = pager.applyFilter(pe); err != nil {
		return nil, err
	}
	conn := &PermissionConnection{Edges: []*PermissionEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = pe.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	pe = pager.applyCursors(pe, after, before)
	pe = pager.applyOrder(pe, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		pe.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := pe.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := pe.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// PermissionOrderField defines the ordering field of Permission.
type PermissionOrderField struct {
	field    string
	toCursor func(*Permission) Cursor
}

// PermissionOrder defines the ordering of Permission.
type PermissionOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     *PermissionOrderField `json:"field"`
}

// DefaultPermissionOrder is the default ordering of Permission.
var DefaultPermissionOrder = &PermissionOrder{
	Direction: OrderDirectionAsc,
	Field: &PermissionOrderField{
		field: permission.FieldID,
		toCursor: func(pe *Permission) Cursor {
			return Cursor{ID: pe.ID}
		},
	},
}

// ToEdge converts Permission into PermissionEdge.
func (pe *Permission) ToEdge(order *PermissionOrder) *PermissionEdge {
	if order == nil {
		order = DefaultPermissionOrder
	}
	return &PermissionEdge{
		Node:   pe,
		Cursor: order.Field.toCursor(pe),
	}
}

// RoleEdge is the edge representation of Role.
type RoleEdge struct {
	Node   *Role  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// RoleConnection is the connection containing edges to Role.
type RoleConnection struct {
	Edges      []*RoleEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *RoleConnection) build(nodes []*Role, pager *rolePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Role
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Role {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Role {
			return nodes[i]
		}
	}
	c.Edges = make([]*RoleEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &RoleEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// RolePaginateOption enables pagination customization.
type RolePaginateOption func(*rolePager) error

// WithRoleOrder configures pagination ordering.
func WithRoleOrder(order *RoleOrder) RolePaginateOption {
	if order == nil {
		order = DefaultRoleOrder
	}
	o := *order
	return func(pager *rolePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultRoleOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithRoleFilter configures pagination filter.
func WithRoleFilter(filter func(*RoleQuery) (*RoleQuery, error)) RolePaginateOption {
	return func(pager *rolePager) error {
		if filter == nil {
			return errors.New("RoleQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type rolePager struct {
	order  *RoleOrder
	filter func(*RoleQuery) (*RoleQuery, error)
}

func newRolePager(opts []RolePaginateOption) (*rolePager, error) {
	pager := &rolePager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultRoleOrder
	}
	return pager, nil
}

func (p *rolePager) applyFilter(query *RoleQuery) (*RoleQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *rolePager) toCursor(r *Role) Cursor {
	return p.order.Field.toCursor(r)
}

func (p *rolePager) applyCursors(query *RoleQuery, after, before *Cursor) *RoleQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultRoleOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *rolePager) applyOrder(query *RoleQuery, reverse bool) *RoleQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultRoleOrder.Field {
		query = query.Order(direction.orderFunc(DefaultRoleOrder.Field.field))
	}
	return query
}

func (p *rolePager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultRoleOrder.Field {
			b.Comma().Ident(DefaultRoleOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Role.
func (r *RoleQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...RolePaginateOption,
) (*RoleConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newRolePager(opts)
	if err != nil {
		return nil, err
	}
	if r, err = pager.applyFilter(r); err != nil {
		return nil, err
	}
	conn := &RoleConnection{Edges: []*RoleEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = r.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	r = pager.applyCursors(r, after, before)
	r = pager.applyOrder(r, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		r.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := r.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// RoleOrderField defines the ordering field of Role.
type RoleOrderField struct {
	field    string
	toCursor func(*Role) Cursor
}

// RoleOrder defines the ordering of Role.
type RoleOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *RoleOrderField `json:"field"`
}

// DefaultRoleOrder is the default ordering of Role.
var DefaultRoleOrder = &RoleOrder{
	Direction: OrderDirectionAsc,
	Field: &RoleOrderField{
		field: role.FieldID,
		toCursor: func(r *Role) Cursor {
			return Cursor{ID: r.ID}
		},
	},
}

// ToEdge converts Role into RoleEdge.
func (r *Role) ToEdge(order *RoleOrder) *RoleEdge {
	if order == nil {
		order = DefaultRoleOrder
	}
	return &RoleEdge{
		Node:   r,
		Cursor: order.Field.toCursor(r),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) *UserQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userPager) applyOrder(query *UserQuery, reverse bool) *UserQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}
	return query
}

func (p *userPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = u.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	u = pager.applyCursors(u, after, before)
	u = pager.applyOrder(u, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
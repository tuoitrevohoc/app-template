// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package ent

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Invoice struct {
		ID           func(childComplexity int) int
		InvoicedTo   func(childComplexity int) int
		LeetCodeLink func(childComplexity int) int
		Title        func(childComplexity int) int
	}

	InvoiceConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	InvoiceEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		Invoices func(childComplexity int, after *Cursor, first *int, before *Cursor, last *int) int
		Node     func(childComplexity int, id int) int
		Nodes    func(childComplexity int, ids []int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Invoice.id":
		if e.complexity.Invoice.ID == nil {
			break
		}

		return e.complexity.Invoice.ID(childComplexity), true

	case "Invoice.invoicedTo":
		if e.complexity.Invoice.InvoicedTo == nil {
			break
		}

		return e.complexity.Invoice.InvoicedTo(childComplexity), true

	case "Invoice.leetCodeLink":
		if e.complexity.Invoice.LeetCodeLink == nil {
			break
		}

		return e.complexity.Invoice.LeetCodeLink(childComplexity), true

	case "Invoice.title":
		if e.complexity.Invoice.Title == nil {
			break
		}

		return e.complexity.Invoice.Title(childComplexity), true

	case "InvoiceConnection.edges":
		if e.complexity.InvoiceConnection.Edges == nil {
			break
		}

		return e.complexity.InvoiceConnection.Edges(childComplexity), true

	case "InvoiceConnection.pageInfo":
		if e.complexity.InvoiceConnection.PageInfo == nil {
			break
		}

		return e.complexity.InvoiceConnection.PageInfo(childComplexity), true

	case "InvoiceConnection.totalCount":
		if e.complexity.InvoiceConnection.TotalCount == nil {
			break
		}

		return e.complexity.InvoiceConnection.TotalCount(childComplexity), true

	case "InvoiceEdge.cursor":
		if e.complexity.InvoiceEdge.Cursor == nil {
			break
		}

		return e.complexity.InvoiceEdge.Cursor(childComplexity), true

	case "InvoiceEdge.node":
		if e.complexity.InvoiceEdge.Node == nil {
			break
		}

		return e.complexity.InvoiceEdge.Node(childComplexity), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Query.invoices":
		if e.complexity.Query.Invoices == nil {
			break
		}

		args, err := ec.field_Query_invoices_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Invoices(childComplexity, args["after"].(*Cursor), args["first"].(*int), args["before"].(*Cursor), args["last"].(*int)), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(int)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]int)), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCreateInvoiceInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "app/schema/ent.graphql", Input: `directive @goField(forceResolver: Boolean, name: String) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION
directive @goModel(model: String, models: [String!]) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION
"""
CreateInvoiceInput is used for create Invoice object.
Input was generated by ent.
"""
input CreateInvoiceInput {
  title: String!
  leetCodeLink: String!
  invoicedTo: String!
}
"""
Define a Relay Cursor type:
https://relay.dev/graphql/connections.htm#sec-Cursor
"""
scalar Cursor
type Invoice implements Node {
  id: ID!
  title: String!
  leetCodeLink: String!
  invoicedTo: String!
}
"""A connection to a list of items."""
type InvoiceConnection {
  """A list of edges."""
  edges: [InvoiceEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type InvoiceEdge {
  """The item at the end of the edge."""
  node: Invoice
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""
An object with an ID.
Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
"""
interface Node @goModel(model: "github.com/tuoitrevohoc/app-template/api/ent.Noder") {
  """The id of the object."""
  id: ID!
}
"""Possible directions in which to order a list of items when provided an ` + "`" + `orderBy` + "`" + ` argument."""
enum OrderDirection {
  """Specifies an ascending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  ASC
  """Specifies a descending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  DESC
}
"""
Information about pagination in a connection.
https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
"""
type PageInfo {
  """When paginating forwards, are there more items?"""
  hasNextPage: Boolean!
  """When paginating backwards, are there more items?"""
  hasPreviousPage: Boolean!
  """When paginating backwards, the cursor to continue."""
  startCursor: Cursor
  """When paginating forwards, the cursor to continue."""
  endCursor: Cursor
}
type Query {
  """Fetches an object given its ID."""
  node(
    """ID of the object."""
    id: ID!
  ): Node
  """Lookup nodes by a list of IDs."""
  nodes(
    """The list of node IDs."""
    ids: [ID!]!
  ): [Node]!
  invoices(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int
  ): InvoiceConnection!
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)

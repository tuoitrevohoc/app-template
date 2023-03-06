# App Template With EntGO & React

## Tech Stacks

- EntGo (https://entgo.io)
- GQLGen (https://gqlgen.com)
- React
- Relay (https://relay.dev)

## How to run

Check out.

0. Run docker compose up

```
docker compose up
```

1. Run API

```
cd api
go mod tidy
go run cmd/main.go
```

2. Run Frontend

```
cd www
yarn install
yarn dev
```

## GraphQL Playground

http://localhost:8080/graphql
https://app-template.8doc.xyz/graphql

## Backend Development Guideline

- Schema definition /api/schema/entity
- GraphQL definition /api/schema/graphql
- GraphQL resolvers /api/pkg/resolvers

### 1. Change the schema

- Change schema in /api/schema/entity
- Run `go generate`

### 2. Change GraphQL Schema

- Change/add schema definition in /api/schema/graphql

Ex. In invoice.graphql, add:

```
  createInvoice(input: CreateInvoiceInput!): Invoice
```

- Run `go generate`

- Add implementation to resolvers

in invoice.resolvers.go

```
func (r *mutationResolver) CreateInvoice(ctx context.Context, input ent.CreateInvoiceInput) (*ent.Invoice, error) {
	return r.Client.Invoice.Create().
		SetInput(input).
		Save(ctx)
}
```

## Front End Development Guideline

### 0. Set up

- Install Relay plugin in VS Code.
- Run relay-compiler

```
yarn relay watch
```

### 1. Execute Query to fetch data

See Home.tsx

### 2. Call Mutation to change data

See InvoiceEditor.tsx

package hgraph

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Query 与 Mutation 的不同在于 并行与串行 执行

// Http Handler
var GraphqlHttpHandler = func(query, mutation interface{}) *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   GraphqlSchema(query, mutation),
		Pretty:   true,
		GraphiQL: true,
	})
}

// Graphql Schema
var GraphqlSchema = func(query, mutation interface{}) *graphql.Schema {
	newSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:        GraphqlObject(query),
		Mutation:     GraphqlObject(mutation),
		Subscription: nil,
		Types:        nil,
		Directives:   nil,
		Extensions:   nil,
	})
	if err != nil {
		// 异常退出
		log.Fatal(err)
	}
	log.Printf("GraphqlSchema Load Success")
	return &newSchema
}

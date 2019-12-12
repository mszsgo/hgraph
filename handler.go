package hgraph

import (
	"github.com/graphql-go/handler"
)

// Http Handler
var GraphqlHttpHandler = func(query, mutation interface{}) *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   GraphqlSchema(query, mutation),
		Pretty:   true,
		GraphiQL: true,
	})
}

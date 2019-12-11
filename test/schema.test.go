package test

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Http Handler
var GraphqlHttpHandler = func() *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   GraphqlSchema(),
		Pretty:   true,
		GraphiQL: true,
	})
}

// Graphql Schema
var GraphqlSchema = func() *graphql.Schema {
	newSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType(),
		// Mutation:     mutationType(),
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

// Graphql Query Type
var queryType = func() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:       "Query",
		Interfaces: nil,
		// Fields:      queryFields,
		IsTypeOf:    nil,
		Description: "查询操作",
	})
}

// Graphql Mutation Type
/*var mutationType = func() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Mutation",
		Interfaces:  nil,
		Fields:      mutationFields,
		IsTypeOf:    nil,
		Description: "更新操作",
	})
}*/

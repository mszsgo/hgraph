package hgraph

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Query 与 Mutation 的不同在于 并行与串行 执行

// Graphql Schema Fields
var (
	queryFields    graphql.Fields
	mutationFields graphql.Fields
)

// 查询字段定义
func MergeQueryFields(args ...graphql.Fields) {
	if queryFields == nil {
		queryFields = make(graphql.Fields)
	}
	for _, arg := range args {
		for k, v := range arg {
			queryFields[k] = v
		}
	}
}

// 更新字段定义
func MergeMutationFields(args ...graphql.Fields) {
	if mutationFields == nil {
		mutationFields = make(graphql.Fields)
	}
	for _, arg := range args {
		for k, v := range arg {
			mutationFields[k] = v
		}
	}
}

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
		Query:        queryType(),
		Mutation:     mutationType(),
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
		Name:        "Query",
		Interfaces:  nil,
		Fields:      queryFields,
		IsTypeOf:    nil,
		Description: "查询操作",
	})
}

// Graphql Mutation Type
var mutationType = func() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Mutation",
		Interfaces:  nil,
		Fields:      mutationFields,
		IsTypeOf:    nil,
		Description: "更新操作",
	})
}

// 分页请求参数定义
var PageArgument = graphql.FieldConfigArgument{
	"limit": &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 20, Description: "一次返回记录行数，默认20"},
	"skip":  &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 0, Description: "跳过记录行数"},
}

// 合并参数
var Argument = func(args ...graphql.FieldConfigArgument) graphql.FieldConfigArgument {
	var newArgument = make(graphql.FieldConfigArgument)
	for _, item := range args {
		for k, v := range item {
			newArgument[k] = v
		}
	}
	return newArgument
}

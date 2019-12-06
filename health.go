package hgraph

import (
	"github.com/graphql-go/graphql"
)

func init() {
	MergeQueryFields(healthQueryFields)
	MergeMutationFields(healthMutationFields)
}

// Graphql Query
var healthQueryFields = graphql.Fields{
	"health": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return "health info", e
		},
		Description: "项目状态查询",
	},
}

// Graphql Mutation
var healthMutationFields = graphql.Fields{
	"w": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return "health info", e
		},
		Description: "w",
	},
}

// Graphql 类型缩写, 减少Schema代码量
package hgraph

import (
	"github.com/graphql-go/graphql"
)

var ArgString = func(description string) *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: description}
}
var ArgNonNullString = func(description string) *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: description}
}
var ArgInt = func(description string) *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: nil, Description: description}
}
var ArgNonNullInt = func(description string) *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int), DefaultValue: nil, Description: description}
}

var NewObject = func(name string, fields graphql.Fields, description string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        name,
		Fields:      fields,
		Description: description,
	})
}

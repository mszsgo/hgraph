package hgraph

import (
	"github.com/graphql-go/graphql"
)

// 扩展Graphql Resolve函数的参数方法
type ResolveParams graphql.ResolveParams

func (ResolveParams) Auth() {

}

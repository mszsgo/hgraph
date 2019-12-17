package test

// 查询类型，生成查询Schema
type Query struct {
	Member *Member `graphql:"!" description:"会员服务"`
	// Version string
}

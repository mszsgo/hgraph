package test

import (
	"github.com/graphql-go/graphql"

	"github.com/mszsgo/hgraph"
)

type Member struct {
	LoginId   string
	Mobile    string
	Email     []string
	Order     []Order
	CreatedAt hgraph.Time
}

// Object 名称，默认用结构体名称
func (m *Member) Name() string {
	return ""
}

// Object 描述
func (*Member) Description() string {
	return "查询操作"
}

// 执行业务逻辑
func (*Member) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		return "", err
	}
}

type MemberArgs struct {
	Uid   string `graphql:"!" description:"用户编号"`
	Skip  int
	Limit int
}

// 会员类型参数
func (*Member) Args() *MemberArgs {
	return &MemberArgs{}
}

type Order struct {
	OrderId string
}

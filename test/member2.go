package test

import (
	"github.com/graphql-go/graphql"

	"github.com/mszsgo/hgraph"
)

type Member2 struct {
	LoginId   string
	Mobile    string
	Email     []string
	Order     []Order2
	CreatedAt hgraph.Time
}

// Object 名称，默认用结构体名称
func (m *Member2) Name() string {
	return "Member222222222222222222222222222"
}

// Object 描述
func (*Member2) Description() string {
	return "查询操作"
}

// 执行业务逻辑
func (*Member2) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		return "", err
	}
}

type MemberArgs2 struct {
	Uid   string `graphql:"!" description:"用户编号"`
	Skip  int
	Limit int
}

// 会员类型参数
func (*Member2) Args() *MemberArgs {
	return &MemberArgs{}
}

type Order2 struct {
	OrderId string
}

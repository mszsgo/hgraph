package test

import (
	"log"
	"time"

	"github.com/graphql-go/graphql"
)

type Member struct {
	LoginId   string
	Mobile    string
	Email     []string
	Order     []Order
	CreatedAt time.Time
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
		log.Print("Resolve------------------")
		i = Member{
			LoginId:   "1132423",
			Mobile:    "12",
			Email:     []string{"1231", "12312"},
			Order:     []Order{{OrderId: "1223423423"}},
			CreatedAt: time.Now(),
		}
		return i, err
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

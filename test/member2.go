package test

import (
	"encoding/json"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/mszsgo/hjson"
)

type Member2 struct {
	LoginId   string
	Mobile    string
	Email     []string
	Order     []Order2
	CreatedAt string
}

// Object 名称，默认用结构体名称
func (m *Member2) Name() string {
	return "Member2"
}

// Object 描述
func (*Member2) Description() string {
	return "查询操作"
}

// 执行业务逻辑
func (*Member2) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args2 *MemberArgs2
		hjson.MapToStruct(p.Args, &args2)
		b, _ := json.Marshal(args2)
		log.Printf(string(b))
		return "", err
	}
}

type MemberArgs2 struct {
	Uid   string `graphql:"!" description:"用户编号123uuu"`
	Order []OrderList
	Skip  int
	Limit int
}

// 会员类型参数
func (*Member2) Args() *MemberArgs2 {
	return &MemberArgs2{}
}

type Order2 struct {
	OrderId string
}

type OrderList struct {
	OrderId string
}

package hgraph

import (
	"time"

	"github.com/graphql-go/graphql"
)

// Graphql 接口字段用户访问权限控制

// 用户身份： toa=运营管理用户   tob=机构商户门店   toc=普通会员用户  tod=对接项目接口
type UserIdentity string

const (
	TOA UserIdentity = "toa"
	TOB UserIdentity = "tob"
	TOC UserIdentity = "toc"
	TOD UserIdentity = "tod"
)

func GetSession(p graphql.ResolveParams) *UserSession {
	// 在入口，根据token查询UserSession，存入上下文中
	return p.Context.Value("session").(*UserSession)
}

type UserSession struct {
	Identity UserIdentity `json:"identity"` // 用户身份
	Token    string       `json:"token"`
	OrgId    string       `json:"orgId"`    // 当前登录用户所属机构编号
	Uid      string       `json:"uid"`      // 登录用户编号
	Expires  time.Time    `json:"expires"`  // 过期时间
	LoginId  string       `json:"loginId"`  // 登录账号
	Nickname string       `json:"nickname"` // 用户昵称
	Avatar   string       `json:"avatar"`   // 用户头像URL

}

func (s *UserSession) Authority(toa, tob, toc, tod func(s *UserSession) error) (err error) {
	switch s.Identity {
	case TOA:
		err = toa(s)
	case TOB:
		err = tob(s)
	case TOC:
		err = toc(s)
	case TOD:
		err = tod(s)
	default:
	}
	return err
}

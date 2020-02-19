package hgraph

import (
	"errors"
	"fmt"
)

// 错误码规则：99*** ，全局使用
var (
	// 通用错误码
	SUCCESS = errors.New("00000:ok")
	FAIL    = errors.New("99999:fail")

	E99101 = errors.New("99101:Query字符串解析错误")

	E99110 = errors.New("99110:Feign服务请求HTTP异常")
	E99111 = errors.New("99111:Feign响应报文状态异常")
	E99112 = errors.New("99112:Feign响应报文读取错误")
)

func Error(userDefinedErr error, args ...interface{}) error {
	if userDefinedErr == nil {
		return errors.New("userDefinedErr is nil")
	}
	e := userDefinedErr.Error()
	errMsg := fmt.Sprintf(e, args...)
	return errors.New(errMsg)
}

package hgraph

import (
	"encoding/json"
	"log"
)

// Graphql 请求JSON Model
type GraphRequestModel struct {
	RequestId     string      `json:"requestId"`
	Token         string      `json:"token"`
	OperationName string      `json:"operationName"`
	Query         string      `json:"query"`
	Variables     interface{} `json:"variables"`
}

// Graphql 响应JSON Model
type GraphResponseModel struct {
	RequestId string                   `json:"requestId"`
	HostTime  string                   `json:"hostTime"`
	Data      map[string]interface{}   `json:"data"`
	Errors    []map[string]interface{} `json:"errors,omitempty"`
}

var ParseGraphqlReuqest = func(b []byte) *GraphRequestModel {
	var model *GraphRequestModel
	err := json.Unmarshal(b, &model)
	if err != nil {
		log.Print(err)
	}
	return model
}

// 网关调用服务
func Gateway(body []byte) []byte {
	reqModel := ParseGraphqlReuqest(body)
	bytes, err := json.Marshal(Feign(reqModel))
	if err != nil {
		log.Print(err)
	}
	return bytes
}

// 调用业务服务
// 解析请求字符串第一级字段作为服务名调用Graphql服务
func Feign(model *GraphRequestModel) *GraphResponseModel {
	services := ParseGraphqlQuery(model.Query)
	return BulkRequest(&GraphRequestModel{
		RequestId:     model.RequestId,
		Token:         model.Token,
		OperationName: model.OperationName,
		Query:         "",
		Variables:     model.Variables,
	}, services)
}

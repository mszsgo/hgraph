package hgraph

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/mszsgo/hjson"
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

// 返回第一条错误内容
func (r *GraphResponseModel) FirstErrorMessage() error {
	if len(r.Errors) > 0 {
		return errors.New(r.Errors[0]["message"].(string))
	}
	return nil
}

// 转为结构体
func (r *GraphResponseModel) ToStruct(serviceName string, output interface{}) {
	s := r.Data[serviceName]
	hjson.MapToStruct(s, output)
}

func ParseGraphqlReuqest(b []byte) (*GraphRequestModel, error) {
	var model *GraphRequestModel
	err := json.Unmarshal(b, &model)
	if err != nil {
		log.Print("Request JSON Parse Error:" + err.Error() + "  ReqJson=" + string(b))
		return nil, errors.New("Request JSON Parse Error")
	}
	return model, nil
}

// 调用业务服务
// 解析请求字符串第一级字段作为服务名调用Graphql服务
func Feign(reqModel *GraphRequestModel) (resModel *GraphResponseModel) {

	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			resModel = &GraphResponseModel{
				RequestId: reqModel.RequestId,
				HostTime:  time.Now().Format(time.RFC3339),
				Data:      map[string]interface{}{},
				Errors:    []map[string]interface{}{{"message": e.Error()}},
			}
		}
	}()

	services := ParseGraphqlQuery(reqModel.Query)
	if len(services) == 0 {
		log.Print("Graphql Query String parse Error hql=" + reqModel.Query)
		panic(errors.New("Graphql Query String parse Error len(services)=0"))
	}
	resModel = BulkRequest(&GraphRequestModel{
		RequestId:     reqModel.RequestId,
		Token:         reqModel.Token,
		OperationName: reqModel.OperationName,
		Query:         "",
		Variables:     reqModel.Variables,
	}, services)
	return resModel
}

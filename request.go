package hgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// 批量请求Graphql服务响应报文，合并响应结果返回
func BulkRequest(reqModel *GraphRequestModel, services map[string]string) *GraphResponseModel {
	resModel := &GraphResponseModel{
		RequestId: reqModel.RequestId,
		HostTime:  time.Now().Format(time.RFC3339),
		Data:      make(map[string]interface{}),
		Errors:    nil,
	}
	for k, v := range services {
		if k == "" {
			continue
		}
		reqModel.Query = v
		serviceResModel, err := PostGraphqlService(k, reqModel)
		if err != nil {
			errorMap := make(map[string]interface{})
			errorMap["message"] = err.Error()
			resModel.Errors = append(resModel.Errors, errorMap)
			continue
		}
		for k2, v2 := range serviceResModel.Data {
			resModel.Data[k2] = v2
		}
		for _, v2 := range serviceResModel.Errors {
			resModel.Errors = append(resModel.Errors, v2)
		}
	}
	return resModel
}

// 根据服务名调用服务
// bytes 为http请求响应body
// message 为错误消息
func PostGraphqlService(name string, requestModel *GraphRequestModel) (resModel *GraphResponseModel, err error) {
	reqBytes, err := json.Marshal(requestModel)
	if err != nil {
		return nil, err
	}
	bytes, err := postRequest(name, reqBytes)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &resModel)
	if err != nil {
		return nil, err
	}
	return
}

// Post 请求，可使用HTTP代理
func postRequest(name string, reqBytes []byte) (resBytes []byte, err error) {
	client := HttpClient()
	resp, err := client.Post(fmt.Sprintf("http://%s/graphql", name), "application/json;charset=utf-8", strings.NewReader(string(reqBytes)))
	if err != nil {
		log.Print("ERR_FEIGN:" + err.Error())
		return nil, E99110
	}
	if resp.StatusCode != 200 {
		log.Printf(fmt.Sprintf("调用服务 %s HTTP响应状态 %d", name, resp.StatusCode))
		return nil, E99111
	}
	resBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ERR_FEIGN:读取响应数据错误" + err.Error())
		return nil, E99112
	}
	return
}

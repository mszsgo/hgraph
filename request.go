package hgraph

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// 批量请求Graphql服务响应报文，合并响应结果返回
func BulkRequest(reqModel *GraphRequestModel, services map[string]string) *GraphResponseModel {
	resModel := &GraphResponseModel{
		RequestId: reqModel.RequestId,
		HostTime:  time.Now().String(),
		Data:      make(map[string]interface{}),
		Errors:    nil,
	}
	for k, v := range services {
		if k == "" {
			continue
		}
		reqModel.Query = v
		serviceResModel := PostGraphql(k, reqModel)
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
func PostGraphql(name string, requestModel *GraphRequestModel) (resModel *GraphResponseModel) {
	resModel = &GraphResponseModel{}
	defer func() {
		if msg := recover(); msg != nil {
			errorObject := make(map[string]interface{})
			errorObject["message"] = msg
			resModel.Errors = []map[string]interface{}{errorObject}
			return
		}
	}()
	reqBytes, err := json.Marshal(requestModel)
	if err != nil {
		panic(err.Error())
		return
	}
	bytes, err := postRequest(name, reqBytes)
	if err != nil {
		panic(err.Error())
		return
	}
	err = json.Unmarshal(bytes, &resModel)
	if err != nil {
		panic(err.Error())
		return
	}
	return
}

// Post 请求，可使用HTTP代理
func postRequest(name string, reqBytes []byte) (resBytes []byte, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			// 服务HTTP代理配置，示例系统环境变量： "MS_HTTP_PROXY=211.152.57.29:39084"
			Proxy: func(request *http.Request) (url *url.URL, err error) {
				msHttpProxy := os.Getenv("MS_HTTP_PROXY")
				if msHttpProxy != "" {
					request.URL.Host = msHttpProxy
				}
				return request.URL, err
			},
		},
		Timeout: 5 * time.Second,
	}
	resp, err := client.Post(fmt.Sprintf("http://%s/graphql", name), "application/json;charset=utf-8", strings.NewReader(string(reqBytes)))
	if err != nil {
		log.Print(err)
		return nil, errors.New("ERR_FEIGN " + err.Error())
	}
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("调用服务 %s HTTP响应状态 %d", name, resp.StatusCode))
	}
	resBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("ERR_FEIGN 读取响应数据错误" + err.Error())
	}
	return resBytes, nil
}

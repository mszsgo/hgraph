package hgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 解析请求字符串第一级字段作为服务名调用Graphql服务
func Graphql(hql string) []byte {
	services := ParseGraphqlMicroService(hql)
	return bulkRequest(services)
}

// 批量请求Graphql服务响应报文
func bulkRequest(services map[string]string) []byte {
	DATA, ERRORS := "data", "errors"

	res := make(map[string]interface{})
	res[DATA] = make(map[string]interface{})
	res[ERRORS] = []map[string]interface{}{}
	for k, v := range services {
		if k == "" {
			continue
		}
		bytes := PostGraphql(k, v)
		var body map[string]interface{}
		err := json.Unmarshal(bytes, body)
		if err != nil {
			log.Print(err)
		}
		for k1, v1 := range res {
			if k1 == DATA {
				for k2, v2 := range body[k1].(map[string]interface{}) {
					v1.(map[string]interface{})[k2] = v2
				}
			}
			if k1 == ERRORS {
				for _, v2 := range body[k1].([]map[string]interface{}) {
					v1 = append(v1.([]map[string]interface{}), v2)
				}
			}
		}
	}
	bytes, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
	}
	return bytes
}

// 根据服务名调用服务
// bytes 为http请求响应body
// message 为错误消息
func PostGraphql(name, hql string) (bytes []byte) {
	defer func() {
		if msg := recover(); msg != nil {
			body := make(map[string]interface{})
			errorObject := make(map[string]interface{})
			errorObject["message"] = msg
			body["errors"] = []map[string]interface{}{errorObject}
			bytes, err := json.Marshal(body)
			if err != nil {
				log.Println("Error Feign convert JSON exception")
			}
			log.Println(bytes)
		}
	}()
	resp, err := http.Post(fmt.Sprintf("http://%s/graphql", name), "application/graphql", strings.NewReader(hql))
	if err != nil {
		log.Print(err)
		panic("ERR_FEIGN " + err.Error())
	}
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("调用服务 %s HTTP响应状态 %d", name, resp.StatusCode))
	}
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("ERR_FEIGN 读取响应数据错误" + err.Error())
	}
	return
}

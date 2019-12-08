package hgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 根据服务名调用服务
// bytes 为http请求响应body
// message 为错误消息
func Call(name, hql string) (bytes []byte) {
	defer func() {
		if msg := recover(); msg != nil {
			body := make(map[string]interface{})
			errorObject := make(map[string]interface{})
			errorObject["message"] = msg
			body["errors"] = []map[string]interface{}{errorObject}
			bytes, err := json.Marshal(body)
			if err != nil {
				log.Println("Error Feign转换JSON异常")
			}
			log.Println(bytes)
		}
	}()
	resp, err := http.Post(fmt.Sprintf("http://%s/api/graphql", name), "text/plain", strings.NewReader(hql))
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

// 合并多个Graphql服务响应报文
func Merge(services map[string]string) []byte {
	res := make(map[string]interface{})
	res["data"] = make(map[string]interface{})
	res["errors"] = []map[string]interface{}{}
	for k, v := range services {
		if k == "" {
			continue
		}
		bytes := Call(k, v)
		var body map[string]interface{}
		err := json.Unmarshal(bytes, body)
		if err != nil {
			log.Print(err)
		}
		for k1, v1 := range res {
			for k2, v2 := range body[k1].(map[string]interface{}) {
				v1.(map[string]interface{})[k2] = v2
			}
		}
	}
	bytes, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
	}
	return bytes
}

// 调用服务
func Graphql(hql string) []byte {
	return Merge(ParseGraphqlMicroService(hql))
}

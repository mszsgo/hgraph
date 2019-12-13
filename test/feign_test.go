package test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/mszsgo/hgraph"
)

func TestParseGraphqlMicroService(t *testing.T) {
	hgraph.ParseGraphqlQuery("query {member(){list(){}   ,  total(){}}  , \r\n points(){list(){},total(){}}\r\n,svf(){list(){\r\n},total(){}}}")
	hgraph.ParseGraphqlQuery("query findMemeber {member{list(){id,name}   ,  total(){}}  }")
	hgraph.ParseGraphqlQuery("mutation {member(){list(){}   ,  total(){}}  }")
}

// 测试服务调用
// query{   member{session(token:"23423"){uid,mobile}},captcha{number{captchaId,base64Image}} }
func TestGraphql(t *testing.T) {
	resModel := hgraph.Feign(&hgraph.GraphRequestModel{
		RequestId:     "13123123123",
		Token:         "13123123123",
		OperationName: "",
		Query:         `query wew{member{session(token:"123123"){uid}},captcha{number{captchaId}} }`,
		Variables:     nil,
	})
	bytes, err := json.Marshal(resModel)
	if err != nil {
		log.Panic(err)
	}
	log.Print(string(bytes))
}

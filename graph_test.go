package hgraph

import (
	"encoding/json"
	"log"
	"testing"
)

func TestParseGraphqlMicroService(t *testing.T) {
	ParseGraphqlQuery("query {member(){list(){}   ,  total(){}}  , \r\n points(){list(){},total(){}}\r\n,svf(){list(){\r\n},total(){}}}")
	ParseGraphqlQuery("query findMemeber {member{list(){id,name}   ,  total(){}}  }")
	ParseGraphqlQuery("mutation {member(){list(){}   ,  total(){}}  }")
}

// 测试服务调用
// query{   member{session(token:"23423"){uid,mobile}},captcha{number{captchaId,base64Image}} }
func TestGraphql(t *testing.T) {
	resModel := Graphql(&GraphRequestModel{
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

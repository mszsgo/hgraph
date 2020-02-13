package hgraph

import (
	"encoding/json"
	"log"
	"time"
)

func responseModel(requestBody []byte) (resModel *GraphResponseModel) {
	reqModel, err := ParseGraphqlReuqest(requestBody)
	if err != nil {
		resModel = &GraphResponseModel{
			RequestId: "",
			HostTime:  time.Now().Format(time.RFC3339),
			Data:      map[string]interface{}{},
			Errors:    []map[string]interface{}{{"message": err.Error()}},
		}
		return
	}
	resModel = Feign(reqModel)
	return
}

// 网关调用服务
func Gateway(requestBody []byte) []byte {
	resModel := responseModel(requestBody)
	responseBody, err := json.Marshal(resModel)
	if err != nil {
		log.Panic("ERROR ->  responseBody, err := json.Marshal(resModel)")
	}
	return responseBody
}

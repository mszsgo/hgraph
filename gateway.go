package hgraph

import (
	"encoding/json"
	"log"
	"time"
)

func jsonRequest(requestBody []byte) (resModel *GraphResponseModel) {
	var reqModel *GraphRequestModel
	err := json.Unmarshal(requestBody, &reqModel)
	if err != nil {
		log.Print("Request JSON Parse Error:" + err.Error() + "  ReqJson=" + string(requestBody))
		resModel = &GraphResponseModel{
			RequestId: UUID(),
			HostTime:  time.Now().Format(time.RFC3339),
			Data:      map[string]interface{}{},
			Errors:    []map[string]interface{}{{"message": "Request JSON Parse Error"}},
		}
		return
	}
	resModel = Feign(reqModel)
	return
}

// 网关调用服务
func Gateway(requestBody []byte) []byte {
	// 待实现：解密

	resModel := jsonRequest(requestBody)
	responseBody, err := json.Marshal(resModel)
	if err != nil {
		log.Printf("ERROR ->  responseBody, err := json.Marshal(resModel)")
	}

	// 待实现：加密
	return responseBody
}

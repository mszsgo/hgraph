package hgraph

import (
	"fmt"
	"log"
)

// 解析Graphql报文，请求到对应服务接口
func ParseGraphqlMicroService(gatewayHql string) map[string]string {
	services := make(map[string]string)
	operation := ""

	lb := 0
	begIndex := 0
	endIndex := 0
	endFlag := false
	skipIndex := 0
	runes := []rune(gatewayHql)
	for i, ch := range runes {
		char := string(ch)
		if char == "{" {
			if operation == "" {
				operation = string(runes[0:i])
			}
			if lb++; lb == 1 {
				begIndex = i
				continue
			}
		}
		if char == "}" {
			if lb--; lb == 1 {
				endFlag = true
				endIndex = i + 1
				continue
			}
		}
		if endFlag {
			skipIndex++
		}
		if endFlag && !(char == " " || char == "," || char == "\n" || char == "\r") {
			endFlag = false
			item := runes[begIndex+1 : endIndex]
			begIndex = endIndex + skipIndex - 2
			skipIndex = 0
			serviceName := ""
			for i, r := range item {
				s := string(r)
				if s == "(" || s == "{" || s == "," {
					serviceName = string(item[0:i])
					break
				}
			}
			v := fmt.Sprintf("%s {%s}", operation, string(item))
			log.Printf("name=【%s】 item=【%s】", serviceName, v)
			services[serviceName] = v
		}
	}
	log.Println(services)
	return services
}

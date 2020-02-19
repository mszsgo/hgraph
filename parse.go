package hgraph

import (
	"fmt"
	"log"
	"strings"
)

// 解析Graphql 字符串，第一级Key作为服务名，value拼接为查询字符串
func ParseGraphqlQuery(query string) (services map[string]string, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ParseGraphqlQuery Error queryString=" + query)
			err = E99101
		}
	}()
	services = make(map[string]string)
	operation := ""

	lb := 0
	begIndex := 0
	endIndex := 0
	endFlag := false
	skipIndex := 0
	runes := []rune(query)
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
					serviceName = strings.TrimSpace(string(item[0:i]))
					break
				}
			}
			v := fmt.Sprintf("%s {%s}", operation, string(item))
			log.Printf("serviceName=%s Graphql: %s", serviceName, v)
			services[serviceName] = v
		}
	}
	return services, nil
}

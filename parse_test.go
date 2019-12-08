package hgraph

import (
	"testing"
)

func TestParseGraphqlMicroService(t *testing.T) {
	ParseGraphqlMicroService("query {member(){list(){}   ,  total(){}}  , \r\n points(){list(){},total(){}}\r\n,svf(){list(){\r\n},total(){}}}")
	ParseGraphqlMicroService("query findMemeber {member(){list(){}   ,  total(){}}  }")
	ParseGraphqlMicroService("mutation {member(){list(){}   ,  total(){}}  }")
}

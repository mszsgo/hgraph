package test

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/mszsgo/hgraph"
)

// 必填 `graphql:"!name"  description:"描述" value:"默认值" deprecationReason:""`

func TestGraphqlType(t *testing.T) {
	object := Query{}
	objectType := reflect.TypeOf(object)
	gtype := hgraph.GraphqlObject(objectType)
	log.Println(gtype.String())
}

func TestType(t *testing.T) {
	fmt.Println(reflect.TypeOf(Query{}))
	fmt.Println(reflect.TypeOf(&Query{}))
	fmt.Println(reflect.ValueOf(Query{}))
	fmt.Println(reflect.ValueOf(&Query{}))
}

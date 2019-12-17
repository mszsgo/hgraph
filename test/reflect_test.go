package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	rt := reflect.TypeOf(&Query{})
	fmt.Println(rt.Kind())
	rt = reflect.TypeOf(Query{})
	fmt.Println(rt.Kind())
	rt = reflect.TypeOf("")
	fmt.Println(rt.Kind())
}

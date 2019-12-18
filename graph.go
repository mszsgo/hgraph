// Graphql 类型缩写, 减少Schema代码量
package hgraph

import (
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
)

// Golang 数据类型对应的 Graphql 数据类型
func scalar(t reflect.Type) graphql.Type {
	switch t.Kind() {
	case reflect.String:
		return graphql.String
	case reflect.Int:
		return graphql.Int
	case reflect.Int64:
		return graphql.Int
	case reflect.Int32:
		return graphql.Int
	case reflect.Int8:
		return graphql.Int
	case reflect.Float32:
		return graphql.Float
	case reflect.Float64:
		return graphql.Float
	case reflect.Bool:
		return graphql.Boolean
	default:
		log.Printf("没有匹配到类型 t.Kind().String() = %s", t.Kind().String())
		return graphql.String
	}
}

type Tag struct {
	Graphql           string
	Description       string
	DeprecationReason string
	DefaultValue      string
}

func getTag(structTag reflect.StructTag) *Tag {
	return &Tag{
		Graphql:           structTag.Get("graphql"),
		Description:       structTag.Get("description"),
		DeprecationReason: structTag.Get("deprecationReason"),
		DefaultValue:      structTag.Get("defaultValue"),
	}
}

func resolve(v reflect.Value) graphql.FieldResolveFn {
	var resolve graphql.FieldResolveFn
	resolveFn, ok := v.Type().MethodByName("Resolve")
	if ok {
		resolve = (resolveFn.Func.Call([]reflect.Value{v})[0]).Interface().(graphql.FieldResolveFn)
	}
	return resolve
}

func description(v reflect.Value) string {
	var d string
	method, ok := v.Type().MethodByName("Description")
	if ok {
		d = (method.Func.Call([]reflect.Value{v})[0]).Interface().(string)
	}
	return d
}

func name(v reflect.Value) string {
	method, ok := v.Type().MethodByName("Name")
	if ok {
		return (method.Func.Call([]reflect.Value{v})[0]).Interface().(string)
	}
	return ""
}

func argsInputType(t reflect.Type) graphql.Type {
	if reflect.Slice == t.Kind() {
		t = t.Elem()
	}
	if reflect.Ptr == t.Kind() {
		t = t.Elem()
	}
	if reflect.Struct != t.Kind() {
		return scalar(t)
	}
	fields := graphql.InputObjectConfigFieldMap{}
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		if structField.Anonymous {
			continue
		}
		var inputType graphql.Type
		if reflect.Slice == structField.Type.Kind() {
			inputType = graphql.NewList(argsInputType(structField.Type.Elem()))
		} else {
			inputType = argsInputType(structField.Type)
		}
		tag := getTag(structField.Tag)
		if tag.Graphql != "" && strings.HasPrefix(tag.Graphql, "!") {
			inputType = graphql.NewNonNull(inputType)
		}
		fields[fieldName(structField)] = &graphql.InputObjectFieldConfig{Type: inputType, DefaultValue: tag.DefaultValue, Description: tag.Description}
	}
	// Graphql 输入对象
	inputObject := graphql.NewInputObject(graphql.InputObjectConfig{
		Name:        t.Name(),
		Fields:      fields,
		Description: "",
	})
	return inputObject
}

func args(v reflect.Value) graphql.FieldConfigArgument {
	argsMethod, ok := v.Type().MethodByName("Args")
	if !ok {
		return nil
	}
	var fieldConfigArgument = graphql.FieldConfigArgument{}
	value := argsMethod.Func.Call([]reflect.Value{v})[0]
	argsType := value.Type().Elem()
	// log.Printf("argsType %s ", argsType.Name())
	for j := 0; j < argsType.NumField(); j++ {
		argField := argsType.Field(j)
		// log.Printf("argField.Name %d=%s ", j, argField.Name)
		if argField.Anonymous {
			continue
		}
		var inputType graphql.Type
		if reflect.Slice == argField.Type.Kind() {
			inputType = graphql.NewList(argsInputType(argField.Type.Elem()))
		} else {
			inputType = argsInputType(argField.Type)
		}
		argTag := getTag(argField.Tag)
		if argTag.Graphql != "" && strings.HasPrefix(argTag.Graphql, "!") {
			inputType = graphql.NewNonNull(inputType)
		}

		// Fields Default Value
		var defValue interface{}
		if argTag.DefaultValue != "" {
			if reflect.Int == argField.Type.Kind() {
				defValue, _ = strconv.Atoi(argTag.DefaultValue)
			}
			if reflect.Int64 == argField.Type.Kind() {
				defValue, _ = strconv.ParseInt(argTag.DefaultValue, 10, 64)
			}
			if reflect.String == argField.Type.Kind() {
				defValue = argTag.DefaultValue
			}
		}
		fieldConfigArgument[fieldName(argField)] = &graphql.ArgumentConfig{
			Type:         inputType,
			DefaultValue: defValue,
			Description:  argTag.Description,
		}
	}
	return fieldConfigArgument
}

func fieldName(field reflect.StructField) string {
	tag := getTag(field.Tag)
	// 如果标签中存在名字则使用标签中定义的字段名，否则使用结构体属性名，首字母小写
	fieldName := strings.ToLower(field.Name[0:1]) + field.Name[1:]
	if tag.Graphql != "" {
		if tag.Graphql[:1] != "!" {
			fieldName = tag.Graphql
		}
		if tag.Graphql[:1] == "!" && len(tag.Graphql) > 1 {
			fieldName = tag.Graphql[1:]
		}
	}
	return fieldName
}

// object := Query{}
// objectType := reflect.TypeOf(object)
func GraphqlType(t reflect.Type) graphql.Type {
	if reflect.Ptr == t.Kind() {
		t = t.Elem()
	}
	if reflect.Struct != t.Kind() {
		return scalar(t)
	}
	if t.Name() == "Time" {
		// 时间格式使用RFC3339标准
		return graphql.DateTime
	}
	fields := graphql.Fields{}
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		if structField.Anonymous {
			continue
		}
		structType := structField.Type

		// Graphql Field Type
		var gtype graphql.Type
		if reflect.Slice == structType.Kind() {
			gtype = graphql.NewList(GraphqlType(structType.Elem()))
		} else {
			gtype = GraphqlType(structType)
		}
		// New Field Type
		var newStructType reflect.Value
		if reflect.Slice == structType.Kind() {
			structType = structType.Elem()
		}
		if reflect.Ptr == structType.Kind() {
			structType = structType.Elem()
		}
		newStructType = reflect.New(structType)
		// Graphql Field Args
		args := args(newStructType)
		// Graphql Field  Resolve
		resolve := resolve(newStructType)
		// Graphql Field Description
		tag := getTag(structField.Tag)
		fields[fieldName(structField)] = &graphql.Field{
			Type:              gtype,
			Args:              args,
			Resolve:           resolve,
			DeprecationReason: tag.DeprecationReason,
			Description:       tag.Description,
		}
	}
	// 调用结构体Name方法，获取Object名字，默认使用结构体名字
	newt := reflect.New(t)
	name := name(newt)
	if name == "" {
		name = t.Name()
	}
	graphqlObject := graphql.NewObject(graphql.ObjectConfig{
		Name:        name,
		Interfaces:  nil,
		Fields:      fields,
		IsTypeOf:    nil,
		Description: description(newt),
	})
	return graphqlObject
}

// i = Query{}
// i = Mutation{}
func GraphqlObject(i interface{}) *graphql.Object {
	return GraphqlType(reflect.TypeOf(i)).(*graphql.Object)
}

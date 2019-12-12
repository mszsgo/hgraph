package test

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/graphql-go/graphql"
)

// 必填 `graphql:"!name"  description:"描述" value:"默认值" deprecationReason:""`

// 查询类型，生成查询Schema
type Query struct {
	Member Member `graphql:"!" description:"会员服务"`
	// Version string
}

type Member struct {
	LoginId string
	// Mobile  string
	// Email   string

	// Order Order
}

// Object 名称，默认用结构体名称
func (*Member) Name() string {
	return ""
}

// Object 描述
func (*Member) Description() string {
	return "查询操作"
}

// 执行业务逻辑
func (*Member) Resolve() interface{} {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		return "", err
	}
}

type MemberArgs struct {
	uid string `graphql:"!" description:"用户编号"`
}

// 会员类型参数
func (*Member) Args() *MemberArgs {
	return &MemberArgs{}
}

type Order struct {
	orderId string
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
	resolveMethod, ok := v.Type().MethodByName("Resolve")
	if ok {
		resolve = (resolveMethod.Func.Call([]reflect.Value{v})[0]).Interface().(graphql.FieldResolveFn)
	}
	return resolve
}

func args(v reflect.Value) graphql.FieldConfigArgument {
	argsMethod, ok := v.Type().MethodByName("Args")
	if !ok {
		return nil
	}
	var fieldConfigArgument = graphql.FieldConfigArgument{}
	value := argsMethod.Func.Call([]reflect.Value{v})[0]
	argsType := value.Type().Elem()
	for j := 0; j < argsType.NumField(); j++ {
		argField := argsType.Field(j)
		inputType := graphqlInputType(argField.Type)
		argTag := getTag(argField.Tag)
		if argTag.Graphql != "" && strings.HasPrefix(argTag.Graphql, "!") {
			inputType = graphql.NewNonNull(inputType)
		}
		fieldConfigArgument[argField.Name] = &graphql.ArgumentConfig{
			Type:         inputType,
			DefaultValue: nil,
			Description:  "",
		}
	}
	return fieldConfigArgument
}

// object := &Query{}
//	objectType := reflect.TypeOf(object).Elem()
func GraphqlObjectType(t reflect.Type) graphql.Type {

	// 结构体类型  对应Graphql Object
	if "struct" == t.Kind().String() {
		fields := graphql.Fields{}
		for i := 0; i < t.NumField(); i++ {
			structField := t.Field(i)
			// 忽略匿名字段
			if structField.Anonymous {
				continue
			}
			// 递归创建类型
			gType := GraphqlObjectType(structField.Type)
			// 解析标签
			tag := getTag(structField.Tag)

			// 调用结构体实例函数方法
			newStructField := reflect.New(structField.Type)

			// 创建字段
			field := &graphql.Field{Type: gType, Args: args(newStructField), Resolve: resolve(newStructField), DeprecationReason: tag.DeprecationReason, Description: tag.Description}
			// 如果标签中存在名字则使用标签中定义的字段名，否则使用结构体属性名，首字母小写
			fieldName := strings.ToLower(structField.Name[0:1]) + structField.Name[1:]
			if tag.Graphql != "" && tag.Graphql[1:] != "" {
				fieldName = tag.Graphql[1:]
			}
			fields[fieldName] = field
		}
		// Graphql对象
		graphqlObject := graphql.NewObject(graphql.ObjectConfig{
			Name:        t.Name(),
			Interfaces:  nil,
			Fields:      fields,
			IsTypeOf:    nil,
			Description: "11",
		})
		return graphqlObject
	}

	// 字符串类型 ，对应Graphql String
	if strings.Contains("string", t.Kind().String()) {
		return graphql.String
	}

	// 数字类型，对应Grahql Int
	if strings.Contains("int,int64", t.Kind().String()) {

		return graphql.Int
	}

	log.Printf("没有匹配到类型 t.Kind().String() = %s", t.Kind().String())
	return nil
}

func graphqlInputType(t reflect.Type) graphql.Type {

	// 结构体类型  对应Graphql Object
	if "struct" == t.Kind().String() {
		fields := graphql.InputObjectConfigFieldMap{}
		for i := 0; i < t.NumField(); i++ {
			structField := t.Field(i)
			// 忽略匿名字段
			if structField.Anonymous {
				continue
			}
			// 递归创建类型
			inputType := graphqlInputType(structField.Type)
			argTagGraphql := structField.Tag.Get("graphql")
			if argTagGraphql != "" && strings.HasPrefix(argTagGraphql, "!") {
				inputType = graphql.NewNonNull(inputType)
			}
			tagDescription := structField.Tag.Get("description")

			// 创建字段
			field := &graphql.InputObjectFieldConfig{Type: inputType, DefaultValue: nil, Description: tagDescription}
			// 如果标签中存在名字则使用标签中定义的字段名，否则使用结构体属性名，首字母小写
			fieldName := strings.ToLower(structField.Name[0:1]) + structField.Name[1:]
			if argTagGraphql != "" && argTagGraphql[1:] != "" {
				fieldName = argTagGraphql[1:]
			}
			fields[fieldName] = field
		}
		// Graphql 输入对象
		inputObject := graphql.NewInputObject(graphql.InputObjectConfig{
			Name:        t.Name(),
			Fields:      fields,
			Description: "",
		})
		return inputObject
	}

	// 字符串类型 ，对应Graphql String
	if strings.Contains("string", t.Kind().String()) {
		return graphql.String
	}

	// 数字类型，对应Grahql Int
	if strings.Contains("int,int64", t.Kind().String()) {
		return graphql.Int
	}
	log.Printf("没有匹配到类型 t.Kind().String() = %s", t.Kind().String())
	return nil
}

func TestGraphqlType(t *testing.T) {
	object := &Query{}
	objectType := reflect.TypeOf(object).Elem()
	gtype := GraphqlObjectType(objectType)
	log.Println(gtype.String())
}

func TestType(t *testing.T) {
	fmt.Println(reflect.TypeOf(Query{}))
	fmt.Println(reflect.TypeOf(&Query{}))
	fmt.Println(reflect.ValueOf(Query{}))
	fmt.Println(reflect.ValueOf(&Query{}))
}

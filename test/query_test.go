package test

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/graphql-go/graphql"
)

// 必填 `graphql:"!name"  description:"描述" value:"默认值" deprecationReason:""`

// 查询类型，生成查询Schema
type Query struct {
	Member  Member `graphql:"!" description:"会员服务"`
	Version string
}

type Member struct {
	LoginId string
	Mobile  string
	Email   string

	Order Order
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
func (*Member) Resolve() func(p graphql.ResolveParams) (i interface{}, err error) {
	return func(p graphql.ResolveParams) (i interface{}, err error) {

		return i, err
	}
}

// 会员类型参数
func (*Member) Args() graphql.FieldConfigArgument {
	args := graphql.FieldConfigArgument{
		"uid": &graphql.ArgumentConfig{
			Type: graphql.String, Description: "用户编号"},
	}
	return args
}

type Order struct {
	orderId string
}

// object := &Query{}
//	objectType := reflect.TypeOf(object).Elem()
func GraphqlType(t reflect.Type) graphql.Type {

	// 结构体类型  对应Graphql Object
	if t.Kind().String() == "struct" {
		fields := graphql.Fields{}
		for i := 0; i < t.NumField(); i++ {
			structField := t.Field(i)
			gType := GraphqlType(structField.Type)
			tagGraphql := structField.Tag.Get("graphql")
			if tagGraphql != "" && strings.HasPrefix(tagGraphql, "!") {
				gType = graphql.NewNonNull(gType)
			}
			tagDescription := structField.Tag.Get("description")
			tagDeprecationReason := structField.Tag.Get("deprecationReason")

			/*method,ok := structField.Type.MethodByName("Resolve")
			if ok {

			}*/

			field := &graphql.Field{Type: gType, Args: nil, Resolve: nil, DeprecationReason: tagDeprecationReason, Description: tagDescription}

			// 如果标签中存在名字则使用标签中定义的字段名，否则使用结构体属性名，首字母小写
			fieldName := strings.ToLower(structField.Name[0:1]) + structField.Name[1:]
			if tagGraphql != "" && tagGraphql[1:] != "" {
				fieldName = tagGraphql[1:]
			}
			fields[fieldName] = field
		}
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
	if t.Kind().String() == "string" {
		return graphql.String
	}

	// 数字类型，对应Grahql Int
	if t.Kind().String() == "int" {
		return graphql.Int
	}

	log.Printf("没有匹配到类型 t.Kind().String() = %s", t.Kind().String())
	return nil
}

func TestGraphqlType(t *testing.T) {
	object := &Query{}
	objectType := reflect.TypeOf(object).Elem()
	gtype := GraphqlType(objectType)
	log.Println(gtype.String())
}

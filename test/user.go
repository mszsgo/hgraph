package test

// graphql:""  value值说明：  使用,号分隔，第一个位字段名，第二个位默认值，第三个为描述

type User struct {
	Id    string `graphql:"id,,用户ID"`
	Name  string
	Age   int     `graphql:"!,,"`
	Class []Class `graphql:"!class"`
}

type Class struct {
	Id   string
	Name string
}

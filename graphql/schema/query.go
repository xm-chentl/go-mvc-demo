package schema

import (
	"github.com/graphql-go/graphql"
)

var queryUsers = graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "hello", nil
	},
}

var querySayHello = graphql.Field{
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type:         graphql.String,
			DefaultValue: 0,
		},
	},
	Type: graphql.String, // 返回值的类型
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "hello " + p.Args["name"].(string), nil
	},
}

// 定义根查询节点
var userQueryObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "user",
	Fields: graphql.Fields{
		"users":    &queryUsers,
		"sayHello": &querySayHello,
	},
})

func GetSchema() (schema graphql.Schema) {
	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    userQueryObject,
		Mutation: nil,
	})

	return
}

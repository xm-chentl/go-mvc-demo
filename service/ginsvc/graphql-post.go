package ginsvc

import (
	graphqlex "github.com/xm-chentl/go-mvc-demo/graphql"
	"github.com/xm-chentl/go-mvc-demo/graphql/mutation"
	"github.com/xm-chentl/go-mvc-demo/graphql/query"
	"github.com/xm-chentl/go-mvc-demo/graphql/sdl"
	"github.com/xm-chentl/go-mvc/ioc"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// NewGraphqlPost
func NewGraphqlPost(g *gin.Engine) {
	schemaString, err := sdl.GetSchemaString()
	if err != nil {
		panic(err)
	}

	mutation := &mutation.Mutation{}
	if err = ioc.Inject(mutation); err != nil {
		panic(err)
	}

	query := &query.Query{}
	if err = ioc.Inject(query); err != nil {
		panic(err)
	}

	schemaSdl := graphql.MustParseSchema(schemaString, &graphqlex.Resolver{
		Mutation: mutation,
		Query:    query,
	})
	g.POST("/graphql", func(c *gin.Context) {
		h := relay.Handler{Schema: schemaSdl}
		h.ServeHTTP(c.Writer, c.Request)
	})
}

// 原生构建
// func nativeBuild(c *gin.Context) {
// 	schemaInst := schema.GetSchema()
// 	h := handler.New(&handler.Config{
// 		Schema:   &schemaInst,
// 		Pretty:   true,
// 		GraphiQL: true,
// 	})
// 	h.ServeHTTP(c.Writer, c.Request)
// }

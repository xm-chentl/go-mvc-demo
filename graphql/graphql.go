package graphal

import (
	"github.com/xm-chentl/go-mvc-demo/graphql/mutation"
	"github.com/xm-chentl/go-mvc-demo/graphql/query"
)

type Resolver struct {
	*mutation.Mutation
	*query.Query
}

package query

import "github.com/xm-chentl/go-code/dbfactory"

type Query struct {
	DbFactory dbfactory.IDbFactory `inject:""`
}

package login

import (
	"github.com/xm-chentl/go-code/dbfactory"
	"github.com/xm-chentl/go-mvc-demo/model/response"
)

type SendVerifiationCodeAPI struct {
	DbFactory dbfactory.IDbFactory `inject:""`

	Phone string `json:"phone"`
}

func (s SendVerifiationCodeAPI) Call() (res interface{}, err error) {
	res = response.EmptyArray
	return
}

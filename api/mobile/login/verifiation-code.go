package login

import "github.com/xm-chentl/go-mvc-demo/model/response"

type VerifiationCodeAPI struct{}

func (s VerifiationCodeAPI) Call() (res interface{}, err error) {
	res = response.EmptyMap
	return
}

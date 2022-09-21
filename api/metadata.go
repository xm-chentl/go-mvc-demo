package api

import (
	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc-demo/api/mobile/login"
	"github.com/xm-chentl/go-mvc-demo/api/mobile/user"
	"github.com/xm-chentl/go-mvc/metadata"
)

// Register is 注册
func Register() {
	metadata.RegisterMap(map[string]mvc.IApi{
		"/mobile/login/account":               &login.AccountAPI{},
		"/mobile/login/send-verifiation-code": &login.SendVerifiationCodeAPI{},
		"/mobile/login/verifiation-code":      &login.VerifiationCodeAPI{},
		"/mobile/user/info":                   &user.InfoAPI{},
	})
}

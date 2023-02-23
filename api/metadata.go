package api

import (
	"github.com/xm-chentl/go-mvc-demo/api/bg/basicdata"
	"github.com/xm-chentl/go-mvc-demo/api/mobile/login"
	"github.com/xm-chentl/go-mvc-demo/api/mobile/user"
	user1 "github.com/xm-chentl/go-mvc-demo/api/user-center/user"

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/metadata"
)

// Register is 注册
func Register() {
	metadata.RegisterMap(map[string]mvc.IApi{
		"/bg/basicdata/import-multi-language":          &basicdata.ImportMultiLanguageAPI{},
		"/mobile/login/account":                        &login.AccountAPI{},
		"/mobile/login/send-verifiation-code":          &login.SendVerifiationCodeAPI{},
		"/mobile/login/verifiation-code":               &login.VerifiationCodeAPI{},
		"/mobile/user/info":                            &user.InfoAPI{},
		"/user-center/user/get_verification_code":      &user1.GetVerificationCodeAPI{},
		"/user-center/user/login":                      &user1.LoginAPI{},
		"/user-center/user/login_by_verification_code": &user1.LoginByVerificationCodeAPI{},
	})
}

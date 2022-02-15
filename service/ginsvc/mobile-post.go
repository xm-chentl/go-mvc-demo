package ginsvc

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/xm-chentl/go-mvc-demo/contract"
	channeltype "github.com/xm-chentl/go-mvc-demo/model/enum/channel-type"
	"github.com/xm-chentl/go-mvc-demo/service/encryptsvc"

	"github.com/gin-gonic/gin"
	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/errorex"
	"github.com/xm-chentl/go-mvc/ginex"
	"github.com/xm-chentl/go-mvc/ioc"
	"github.com/xm-chentl/go-mvc/metadata"
	"github.com/xm-chentl/go-mvc/verify/validator"
)

var encryptKey = "#wwwlebansoftcom#2019!@#$%^&*()_"

type mobileBody struct {
	Value string `json:"lb"`
}

type mobileHeader struct {
	DebugUID      string            `header:"Debug-UID"` // DEBUG模式下的用户UID
	DeviceChannel channeltype.Value `header:"H-C"`       // 渠道类型(channel-type)
	DeviceUnique  string            `header:"H-U"`       // 设备唯一编号
	Mode          string            `header:"H-M"`       // 请求模式
	Token         string            `header:"H-T"`       // 请求token
}

type mobileURI struct {
	Targat   string `uri:"target"`
	EndPoint string `uri:"endpoint"`
}

func NewMobilePost() ginex.Option {
	return func(g *gin.Engine) {
		verifyInst := validator.New()
		g.POST("/mobile/:target/:endpoint", func(ctx *gin.Context) {
			defer func() {
				if recoverErr := recover(); recoverErr != nil {
					fmt.Println("err :", recoverErr)
					ctx.JSON(
						http.StatusOK,
						map[string]interface{}{
							"err":  errorex.ServerErr,
							"data": recoverErr,
						},
					)
				}
			}()

			var err error
			var mobileURI mobileURI
			if err = ctx.ShouldBindUri(&mobileURI); err != nil {
				panic(err)
			}

			var mobileHeader mobileHeader
			if err = ctx.ShouldBindHeader(&mobileHeader); err != nil {
				panic(err)
			}
			code := fmt.Sprintf("/mobile/%s/%s", mobileURI.Targat, mobileURI.EndPoint)
			if ok := metadata.Has(code); !ok {
				panic("api is not exist")
			}

			api := metadata.Get(code)
			newApiInstance := reflect.New(reflect.TypeOf(api).Elem()).Interface()
			isDebug := false
			if mobileHeader.Mode == "DEBUG" {
				// todo: 前期方便调试
				isDebug = true
				if err = ctx.ShouldBindJSON(&newApiInstance); err != nil {
					panic(err)
				}
			} else {
				var mobileBody mobileBody
				if err = ctx.BindJSON(&mobileBody); err != nil {
					panic(err)
				}

				// todo: 如果前端使用未编码的文本，则这里不需要处理
				valueBytes, _ := base64.StdEncoding.DecodeString(mobileBody.Value)
				bodyByte, err := encryptsvc.Decrypt(valueBytes, []byte(encryptKey))
				if err != nil {
					panic(err)
				}
				if err = json.Unmarshal(bodyByte, &newApiInstance); err != nil {
					panic(err)
				}
			}
			if ok := verifyInst.Execute(newApiInstance); !ok {
				panic("参数不正确")
			}

			err = ioc.Inject(newApiInstance, func(field reflect.StructField) interface{} {
				if strings.Contains(field.Type.Name(), "IRoute") {
					return ginex.NewRoute(ctx)
				}

				return nil
			})
			if err != nil {
				panic(err)
			}

			var ok bool
			var mobileDevice contract.IMobileDevice
			mobileDevice, ok = newApiInstance.(contract.IMobileDevice)
			if ok {
				mobileDevice.SetDevice(mobileHeader.DeviceChannel, mobileHeader.DeviceUnique)
			}

			var sessionData contract.ISessionData
			sessionData, ok = newApiInstance.(contract.ISessionData)
			if ok {
				if isDebug {
					sessionData.SetSession(mobileHeader.DebugUID)
				} else {
					// todo: 根据token获取用户信息并赋值全sessionData.SetSession
				}
			}

			var apiInstance mvc.IApi
			apiInstance, ok = newApiInstance.(mvc.IApi)
			if !ok {
				panic("api is not mvc.IApi")
			}

			resp, err := apiInstance.Call()
			if err != nil {
				panic(err)
			}
			ctx.JSON(
				http.StatusOK,
				map[string]interface{}{
					"err":  0,
					"data": resp,
				},
			)
		})
	}
}

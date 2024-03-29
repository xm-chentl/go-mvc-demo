package ginsvc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xm-chentl/go-mvc/context"
	"github.com/xm-chentl/go-mvc/enum"
	"github.com/xm-chentl/go-mvc/errorex"
	"github.com/xm-chentl/go-mvc/ginex"
	"github.com/xm-chentl/go-mvc/handler"
	"github.com/xm-chentl/go-mvc/verify/validator"
	"github.com/xm-chentl/gocore/frame"
)

func NewBgPost() frame.GinOption {
	return func(ginInst *gin.Engine) {
		verifyInst := validator.New()
		handler := handler.Default()
		ginInst.POST("/bg/:server/:action", func(ctx *gin.Context) {
			defer func() {
				if recoverErr := recover(); recoverErr != nil {
					ctx.JSON(
						http.StatusOK,
						map[string]interface{}{
							"err":  errorex.ServerErr,
							"data": map[string]interface{}{},
						},
					)
				}
			}()

			// route := fmt.Sprintf("/bg/%s/%s", ctx.Param("server"), ctx.Param("action"))
			// if ok := metadata.Has(route); !ok {

			// }

			c := context.New()
			c.Set(enum.Code, fmt.Sprintf("/bg/%s/%s", ctx.Param("server"), ctx.Param("action")))
			c.Set(enum.CTX, ginex.NewRoute(ctx))
			c.Set(enum.Verify, verifyInst)
			handler.Execute(c)
		})
	}
}

package ginsvc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xm-chentl/gocore/frame"
	"github.com/xuri/excelize/v2"
)

func UseOrigin() frame.GinOption {
	return func(g *gin.Engine) {
		g.Use(func(ctx *gin.Context) {
			origin := ctx.Request.Header.Get("Origin") //请求头部
			if origin != "" {
				//接收客户端发送的origin （重要！）
				ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				//服务器支持的所有跨域请求的方法
				ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
				//允许跨域设置可以返回其他子段，可以自定义字段
				ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
				// 允许浏览器（客户端）可以解析的头部 （重要）
				ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
				//设置缓存时间
				ctx.Header("Access-Control-Max-Age", "172800")
				//允许客户端传递校验信息比如 cookie (重要)
				ctx.Header("Access-Control-Allow-Credentials", "true")
			}
			ctx.Next()
		})
	}
}

func NewUPload() frame.GinOption {
	return func(g *gin.Engine) {
		g.POST("/upload", func(ctx *gin.Context) {
			var err error
			defer func() {
				respData := map[string]interface{}{
					"code": 0,
					"msg":  "",
				}
				if err != nil {
					respData["msg"] = err.Error()
				}
				ctx.JSON(http.StatusOK, respData)
			}()

			f, _, err := ctx.Request.FormFile("file")
			if err != nil {
				return
			}
			defer f.Close()

			// out, err := os.Create(fHeader.Filename)
			// if err != nil {
			// 	return
			// }

			// _, err = io.Copy(out, f)
			// if err != nil {
			// 	return
			// }

			ef, err := excelize.OpenReader(f)
			if err != nil {
				return
			}
			defer ef.Close()

			rows, err := ef.GetRows("sheet1")
			if err != nil {
				return
			}

			for _, cols := range rows {
				for _, col := range cols {
					fmt.Println("col >>> ", col)
				}
			}
		})
	}
}

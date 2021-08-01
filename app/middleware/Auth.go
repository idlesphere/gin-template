package middleware

import (
	ctl "gin-template/app/controller"
	"gin-template/app/pkg/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.JSON(200, &ctl.ResponseBody{Code: 401, Msg: "No token in header", Data: ""})
			ctx.Abort()
			return
		}

		if _, err := auth.New().Parse(token); err != nil {
			ctx.JSON(200, &ctl.ResponseBody{Code: 401, Msg: "Invalid Token", Data: ""})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

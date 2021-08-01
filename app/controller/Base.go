package controller

import (
	"github.com/gin-gonic/gin"
)

var (
	msg_dict = map[int]string{
		200: "Success",
		201: "Login success",
		202: "Deleted",
		400: "Invalid params",
		401: "Unauthorized",
		403: "Forbidden",
		404: "Not Found",
		405: "Method invalid",
		408: "Request timeout",
		500: "Internal Server Error",
		502: "Bad Gateway",
		511: "Token expired",
	}
)

type ResponseBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BaseController struct {
}

func (ctl BaseController) Response(c *gin.Context, code int, data interface{}) {
	var msg string
	if _, ok := msg_dict[code]; ok {
		msg = msg_dict[code]
	} else {
		msg = "Unknown"
	}

	c.JSON(200, &ResponseBody{Code: code, Msg: msg, Data: data})
}

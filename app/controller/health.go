package controller

import (
	"github.com/gin-gonic/gin"
)

type HealthController struct{
	BaseController
}


func (ctl *HealthController) Run(ctx *gin.Context) {
	ctl.Response(ctx, 200, "pong")
}


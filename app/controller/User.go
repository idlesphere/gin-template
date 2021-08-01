package controller

import (
	"gin-template/app/dto"
	"gin-template/app/pkg/auth"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (ctl *UserController) Run(ctx *gin.Context) {

}

func (ctl *UserController) Login(ctx *gin.Context) {
	var data dto.Login

	// ShouldBind request body data to var data and check if all details are provided
	if ctx.ShouldBind(&data) != nil {
		ctl.Response(ctx, 400, "")
		return
	}

	// TODO: if okta success

	token, err := auth.New().Generate("Gingod")
	//token, err := auth.New().Generate(data.username)

	// If we fail to generate token for access
	if err != nil {
		ctl.Response(ctx, 403, "")
		return
	}

	// Response with header token
	ctx.Header("token", token)
	ctl.Response(ctx, 201, "")
}

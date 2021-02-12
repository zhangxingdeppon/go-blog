package api

import (
	"blog/internal/service"
	"github.com/gin-gonic/gin"
	"blog/pkg/errorcode"
	"blog/global"
	"blog/pkg/app"
)

func GetAuth(c *gin.Context)  {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
		errRsp := errorcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	} 

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Fatalf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errorcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Fatalf("app.GernerateToken err: %v", err)
		response.ToErrorResponse(errorcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
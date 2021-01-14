package v1

import (
	"blog/global"
	"github.com/gin-gonic/gin"
"blog/pkg/app"
"blog/pkg/errorcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {

}
// @Summary Get Multi lable
// @Produce json
// @Param name query string false "label name" maxlength(100)
// @Param state query int false "status" Enums(0,1) default(1)
// @Param page query int false "page code"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errorcode.Error  "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
param := struct {
	Name string `form:"name" binding:"max=100"`
    State uint8 `form:"state, default=1" binding:"oneof=0 1"`
}{}
response := app.NewResponse(c)
valid,errs:=app.BindAndValid(c, &param)
if !valid {
	global.Logger.Fatalf("app.BindAndValid errs: %v", errs)
	errRsp := errorcode.InvalidParams.WithDetails(errs.Errors()...)
	response.ToErrorResponse(errRsp)
	return
}
response.ToResponse(gin.H{})
return

}
// @Summary create lable
// @Produce json
// @Param name body string true "label name" minlength(3) maxlength(100)
// @Param state body int false "status" Enums(0,1) default(1)
// @Param create_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errorcode.Error  "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {

}
// @Summary update lable
// @Produce json
// @Param id path int true "标签ID"
// @Param name body string true "label name" minlength(3) maxlength(100)
// @Param state body int false "status" Enums(0,1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errorcode.Error  "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {

}
// @Summary delete  lable
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errorcode.Error  "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {

}

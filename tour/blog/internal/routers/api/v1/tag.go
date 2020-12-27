package v1

import "github.com/gin-gonic/gin"

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
func (t Tag) List(c *gin.Context) {

}
func (t Tag) Create(c *gin.Context) {

}
func (t Tag) Update(c *gin.Context) {

}
func (t Tag) Delete(c *gin.Context) {

}

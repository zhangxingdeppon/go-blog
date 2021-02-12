package routers

import (
	"blog/global"
	//"blog/pkg/upload"
	v1 "blog/internal/routers/api/v1"
	"blog/internal/routers/api"
	_ "blog/docs"
ginSwagger "github.com/swaggo/gin-swagger"
"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/gin-gonic/gin"
	"blog/internal/middleware"
	"net/http"
	
)

func NewRouter() *gin.Engine {
	r := gin.New()
    r.Use(middleware.Tracing())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
    r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
    r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		// 创建标签
		apiv1.POST("/tags", tag.Create)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", tag.Delete)
		// 更新指定标签
		apiv1.PUT("/tags/:id", tag.Update)
		// 获取标签列表
		apiv1.GET("/tags", tag.List)

		// 创建文章
		apiv1.POST("/articles", article.Create)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", article.Delete)
		// 更新指定文章
		apiv1.PUT("/articles/:id", article.Update)
		// 获取指定文章
		apiv1.GET("/articles/:id", article.Get)
		// 获取文章列表
		apiv1.GET("/articles", article.List)
	}

	return r
}

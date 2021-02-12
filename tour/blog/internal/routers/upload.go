package routers
import (
	"blog/pkg/upload"
	"blog/pkg/convert"
	"blog/pkg/app"
	"github.com/gin-gonic/gin"
	"blog/pkg/errorcode"
	"blog/internal/service"
	"blog/global"
)
type Upload struct {}

func NewUpload() Upload  {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context)  {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errRsp := errorcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return 
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errorcode.InvalidParams)
		return
	}
	svc :=  service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType),file, fileHeader)
	if err != nil {
		global.Logger.Fatalf("svc.UploadFile err: %v", err)
		errRsp := errorcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return 
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
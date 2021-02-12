package errorcode

var (
	ErrorGetTagListFail = NewError(20010001,"获取标签失败")
	ErrorCreateTagFail = NewError(20010002,"创建标签失败")
	ErrorUpdateFail = NewError(20010003,"更新标签失败")
	ErrorDeletedTagFail = NewError(20010004,"删除标签失败")
	ErrorCountTagFail = NewError(20010005,"统计标签失败")
	ErrorUploadFileFail = NewError(20030001,"上传文件失败")
)
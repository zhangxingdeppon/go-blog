package global

import (
	"blog/pkg/logger"
	"blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppsSettingS
	DatabaseSetting *setting.DateBaseSettingS
	Logger          *logger.Logger
)

package model

import (
	"time"
	"blog/global"
	"blog/pkg/setting"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Model 公共部分
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy   string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreateOn   uint32 `json:"create_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DateBaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	db, err := gorm.Open(databaseSetting.DBType, s)
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete",deleteCallBack)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}


func updateTimeStampForCreateCallback(scope *gorm.Scope)  {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn");ok {
			if createTimeField.IsBlank{
				_=createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn");ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}


func updateTimeStampForUpdateCallback(scope *gorm.Scope)  {
	if _,ok:=scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn",time.Now().Unix())
	}
}

func deleteCallBack(scope *gorm.Scope)  {
	if !scope.HasError(){
		var extraOption string
		if str,ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deleteOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDeleted := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDeleted {
			now :=	time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v, %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
                addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string  {
	if str !=""{
			return " " + str
	}
	return ""
}
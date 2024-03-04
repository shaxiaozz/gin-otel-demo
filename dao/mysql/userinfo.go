package mysql

import (
	"errors"
	"gin-otel-demo/db"
	"gin-otel-demo/model"
	"github.com/wonderivan/logger"
)

var UserInfo userInfo

type userInfo struct {
}

// 获取用户信息数据表所有数据
func (u *userInfo) GetListFunc() (data []*model.UserInfo, rowsAffected int64, err error) {
	result := db.GORM.Find(&data)
	if result.Error != nil && result.Error.Error() != "record not found" {
		logger.Error("获取用户信息数据表所有数据失败: " + result.Error.Error())
		return nil, 0, errors.New("获取用户信息数据表所有数据失败: " + result.Error.Error())
	}
	if result.RowsAffected == 0 {
		logger.Info("获取用户信息数据表所有数据不存在...")
		return nil, 0, nil
	}
	return data, result.RowsAffected, nil
}

// 新增数据
func (u *userInfo) InsertFunc(insertData *model.UserInfo) (err error) {
	result := db.GORM.Create(&insertData)
	if result.Error != nil && result.Error.Error() != "record not found" {
		logger.Error("新增用户信息数据表数据失败: " + result.Error.Error())
		return errors.New("新增用户信息数据表数据失败: " + result.Error.Error())
	}
	return nil
}

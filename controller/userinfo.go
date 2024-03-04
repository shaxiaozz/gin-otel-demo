package controller

import (
	"gin-otel-demo/dao/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var UserInfo userInfo

type userInfo struct {
}

func (u *userInfo) GetListHandler(ctx *gin.Context) {
	data, total, err := mysql.UserInfo.GetListFunc()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  err.Error(),
			"data": nil,
		})
	}

	// 返回数据
	ctx.JSON(http.StatusOK, gin.H{
		"code":  1000,
		"msg":   "用户信息数据查询成功",
		"data":  data,
		"total": total,
	})
}

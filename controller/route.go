package controller

import (
	"github.com/gin-gonic/gin"
)

// 实例化router结构体，可以使用该对象点出首字母大写的方法（跨包调用）
var Router router

// 声明一个router结构体
type router struct{}

// 初始化路由规则
func (r *router) InitApiRouter(router *gin.Engine) {
	router.GET("/api/userinfo", UserInfo.GetListHandler). // 获取用户详情
								POST("/api/access_token", AccessToken.AddHandler) // 生成AccessToken
}

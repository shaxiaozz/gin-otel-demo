package main

import (
	"gin-otel-demo/controller"
	"gin-otel-demo/db"
	"gin-otel-demo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitMysql()                          // 初始化mysql数据库
	db.InitRedisDB()                        // 初始化redis数据库
	service.UserInfo.InitUserInfoDataFunc() // 初始化mysql userinfo表数据

	r := gin.Default()
	// 跨包调用router的初始化方法
	controller.Router.InitApiRouter(r)
	// 启动gin server
	r.Run("0.0.0.0:9090")
}

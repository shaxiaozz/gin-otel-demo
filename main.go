package main

import (
	"gin-otel-demo/controller"
	"gin-otel-demo/db"
	"gin-otel-demo/service"
	"gin-otel-demo/trace"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	db.InitMysql()                          // 初始化mysql数据库
	db.InitRedisDB()                        // 初始化redis数据库
	service.UserInfo.InitUserInfoDataFunc() // 初始化mysql userinfo表数据

	// start otel tracing
	if shutdown := trace.Trace.RetryInitTracer(); shutdown != nil {
		defer shutdown()
	}
	r := gin.Default()
	r.Use(otelgin.Middleware("gin-otel-demo")) // 注入 OpenTelemetry 中间件
	controller.Router.InitApiRouter(r)         // 跨包调用router的初始化方法
	r.Run("0.0.0.0:9090")                      // 启动gin server
}

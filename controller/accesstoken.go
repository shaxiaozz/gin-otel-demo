package controller

import (
	"gin-otel-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var AccessToken accessToken

type accessToken struct {
}

func (a *accessToken) AddHandler(ctx *gin.Context) {
	accesstoken, err := service.AccessToken.RandomAccessToken(10)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  err.Error(),
			"data": nil,
		})
	}

	// 返回数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1000,
		"msg":  "accessToken生成成功...",
		"data": accesstoken,
	})
}

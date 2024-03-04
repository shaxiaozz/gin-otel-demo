package service

import (
	"gin-otel-demo/dao/redis"
	"math/rand"
	"time"
)

var AccessToken accessToken

type accessToken struct {
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// 生成固定长度的随机accesstoken
func (a *accessToken) RandomAccessToken(n int, allowedChars ...[]rune) (accessToken string, err error) {
	rand.Seed(time.Now().UnixMicro())
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	if err := redis.AccessToken.GetAccessTokenFunc(string(b)); err != nil {
		return "", err
	}
	return string(b), nil
}

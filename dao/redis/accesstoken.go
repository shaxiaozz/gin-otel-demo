package redis

import (
	"context"
	"gin-otel-demo/db"
	"github.com/redis/go-redis/v9"
	"github.com/wonderivan/logger"
	"time"
)

var AccessToken accessToken

type accessToken struct {
}

// 写入Redis Key
func (a *accessToken) GetAccessTokenFunc(keyName string) (err error) {
	ctx := context.Background()
	// 从Redis获取token
	token, err := db.RDB.Get(ctx, keyName).Result()
	switch {
	case err == redis.Nil:
		logger.Info("Redis Key" + keyName + " 不存在~")
		// 将新token存储到Redis并设置过期时间
		err = db.RDB.Set(ctx, keyName, "1", time.Duration(6)*time.Hour).Err()
		if err != nil {
			return err
		}
		return nil
	case err != nil:
		return err
	case token == "":
		logger.Info("Redis Key" + keyName + "值为空")
		// 将新token存储到Redis并设置过期时间
		err = db.RDB.Set(ctx, keyName, "1", time.Duration(6)*time.Hour).Err()
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

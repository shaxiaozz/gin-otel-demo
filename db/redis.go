package db

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/wonderivan/logger"
	"os"
)

var (
	RDB *redis.Client
)

// 初始化Redis连接
func InitRedisDB() (err error) {
	ctx := context.Background()
	// 判断redis 架构
	redisModel, ok := os.LookupEnv("REDIS_MODEL")
	if !ok {
		redisModel = `standalone`
		os.Setenv("REDIS_MODEL", redisModel)
	}
	redisHost, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		redisHost = `10.82.69.69:6379`
		os.Setenv("REDIS_HOST", redisHost)
	}
	redisPassword, ok := os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		redisPassword = `d8m413WGF9CV`
		os.Setenv("REDIS_PASSWORD", redisPassword)
	}

	// 判断redis 架构
	switch redisModel {
	case "standalone":
		RDB = redis.NewClient(&redis.Options{
			Addr:     redisHost,
			Password: redisPassword,
			DB:       0,
		})
		_, err := RDB.Ping(ctx).Result()
		if err != nil {
			logger.Error("初始化Redis连接失败：" + err.Error())
			return errors.New("初始化Redis连接失败：" + err.Error())
		}
		logger.Info("初始化Redis连接成功~")
	default:
		logger.Error(`请在配置文件中输入正确的Redis架构模式，可选值为：standalone,cluster,sntinel`)
		return errors.New(`请在配置文件中输入正确的Redis架构模式，可选值为：standalone,cluster,sntinel`)
	}
	return nil
}

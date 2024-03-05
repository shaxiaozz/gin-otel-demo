package db

import (
	"gin-otel-demo/model"
	"github.com/wonderivan/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	mysql_logger "gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"
	"time"
)

var (
	isInit bool
	GORM   *gorm.DB
	err    error
)

// 初始化mysql，并建立连接
func InitMysql() {
	// 判断是否已经初始化了
	if isInit {
		return
	}

	// 日志配置（打印慢 SQL 和错误）
	newLogger := mysql_logger.New(
		logrus.NewWriter(),
		mysql_logger.Config{
			SlowThreshold:             time.Second,         // 慢 SQL 阈值
			LogLevel:                  mysql_logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,                // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,               // 禁用彩色打印
		},
	)

	// 组装连接配置
	mySqlUsername, ok := os.LookupEnv("MYSQL_USERNAME")
	if !ok {
		mySqlUsername = "root"
		os.Setenv("MYSQL_USERNAME", mySqlUsername)
	}
	mySqlPassword, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		mySqlPassword = `MK5i^6@S1d`
		os.Setenv("MYSQL_PASSWORD", mySqlPassword)
	}
	mySqlHost, ok := os.LookupEnv("MYSQL_HOST")
	if !ok {
		mySqlHost = "10.82.69.76:3306"
		os.Setenv("OTEL_SERVICE_NAME", mySqlHost)
	}
	mySqlDbName, ok := os.LookupEnv("MYSQL_DB_NAME")
	if !ok {
		mySqlDbName = "gin-otel-demo"
		os.Setenv("MYSQL_DB_NAME", mySqlDbName)
	}
	dsn := mySqlUsername + ":" + mySqlPassword + "@tcp(" + mySqlHost + ")/" + mySqlDbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	GORM, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("MySQL数据库连接失败: " + err.Error())
	}

	// 设置跟踪和指标
	if err := GORM.Use(tracing.NewPlugin()); err != nil {
		panic("MySQL数据库设置OTEL追踪失败: " + err.Error())
	}

	// 数据库连接池配置
	// 获取通用数据库对象
	sqlDB, err := GORM.DB()
	if err != nil {
		panic("获取通用MySQL数据库对象失败: " + err.Error())
	}
	// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于DbMaxIdeleConns常量，超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(20)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(30)
	// 设置连接可复用最大时间
	sqlDB.SetConnMaxLifetime(30 * time.Second)

	// 创建业务表
	user := new(model.UserInfo)
	if err = GORM.AutoMigrate(user); err != nil {
		logger.Error("MySQL数据库 userinfo数据表创建失败: " + err.Error())
		return
	}
	logger.Info("MySQL数据库连接成功~")
	logger.Info("MySQL数据库 userinfo数据表创建成功~")
	isInit = true
}

// 关闭数据库连接
func CloseMysql() (err error) {
	sqlDB, err := GORM.DB()
	return sqlDB.Close()
}

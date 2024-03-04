package service

import (
	"fmt"
	"gin-otel-demo/dao/mysql"
	"gin-otel-demo/model"
	"github.com/wonderivan/logger"
	"math/rand"
	"time"
)

var UserInfo userInfo

type userInfo struct {
}

// 生成随机用户名
func generateUsername() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 生成随机国家
func generateCountry() string {
	countries := []string{"USA", "China", "India", "Russia", "Brazil", "Japan", "Germany", "UK", "France", "Italy"}
	return countries[rand.Intn(len(countries))]
}

// 生成随机城市
func generateCity() string {
	cities := []string{"New York", "Beijing", "Tokyo", "Moscow", "São Paulo", "London", "Paris", "Berlin", "Delhi", "Rome"}
	return cities[rand.Intn(len(cities))]
}

// 生成随机职业
func generateProfession() string {
	professions := []string{"Engineer", "Doctor", "Teacher", "Programmer", "Artist", "Lawyer", "Chef", "Nurse", "Scientist", "Writer"}
	return professions[rand.Intn(len(professions))]
}

// 初始化用户信息表数据
func (u *userInfo) InitUserInfoDataFunc() (err error) {
	// 查询是否存在数据
	_, total, err := mysql.UserInfo.GetListFunc()
	if total != 0 {
		logger.Info("MySQL数据库 userinfo表数据已存在数据，因此不进行初始化操作，如需重新初始化请删除MySQL数据库中的userinfo表...")
		return nil
	}

	// 初始化数据
	logger.Info("即将初始化MySQL数据库 userinfo表数据...")
	rand.Seed(time.Now().UnixNano())
	numRecords := 5000
	for i := 0; i < numRecords; i++ {
		insertData := &model.UserInfo{
			UserName:   generateUsername(),
			Country:    generateCountry(),
			City:       generateCity(),
			Profession: generateProfession(),
		}
		mysql.UserInfo.InsertFunc(insertData)
	}
	logger.Info("MySQL数据库 userinfo表数据已初始化" + fmt.Sprint(numRecords) + "条数据...")
	return nil
}

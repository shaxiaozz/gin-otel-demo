package model

import "time"

// 用户信息数据表
type UserInfo struct {
	ID         uint       `json:"id" gorm:"primaryKey" gorm:"column:id"`                // 自增ID
	CreatedAt  *time.Time `json:"created_at" gorm:"type:datetime;column:created_at"`    // 创建时间
	UpdatedAt  *time.Time `json:"updated_at" gorm:"type:datetime;column:updated_at"`    // 更新时间
	DeletedAt  *time.Time `json:"deleted_at" gorm:"type:datetime;column:deleted_at"`    // 删除时间
	UserName   string     `json:"user_name" gorm:"type:varchar(64);column:user_name"`   // 用户名
	Country    string     `json:"country" gorm:"type:varchar(64);column:country"`       // 国家
	City       string     `json:"city" gorm:"type:varchar(64);column:city"`             // 城市
	Profession string     `json:"profession" gorm:"type:varchar(64);column:profession"` // 职业
}

// / 定义TableName方法，返回mysql表名，以此来定义mysql中的表名
func (*UserInfo) TableName() string {
	return "userinfo"
}

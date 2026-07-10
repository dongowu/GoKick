package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// TODO: 添加业务字段
}

func (User) TableName() string {
	return "users"
}

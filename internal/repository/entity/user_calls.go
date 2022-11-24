package entity

import "time"

type UserCalls struct {
	Id        int64      `gorm:"column=id;PRIMARY" json:"id,omitempty"`
	UserName  string     `gorm:"column=user_name;index" json:"user_name,omitempty"`
	Duration  int64      `gorm:"column=duration;" json:"duration"`
	Block     int64      `gorm:"column=block" json:"block"`
	CreatedAt time.Time  `gorm:"column=created_at;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"column=updated_at;autoUpdateTime" json:"updated_at,omitempty"`
}

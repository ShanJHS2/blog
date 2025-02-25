package models

import (
	"gorm.io/gorm" 
)


type Owner struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null"`//查询一下是否会存在中文不兼容的情况
	Motto string `gorm:"type:varchar(255);not null"`
	AvatarURL string `gorm:"type:varchar(255);not null"`
	GithubLink string `gorm:"type:varchar(255);not null"`
	LeetcodeLink string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"type:varchar(255);not null"`
	Gitee string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

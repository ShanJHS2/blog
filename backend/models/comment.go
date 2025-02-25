package models

import "gorm.io/gorm"

type Comment struct {
    gorm.Model
    BlogID  uint   `gorm:"not null" json:"blogID"`
    Username  string `gorm:"not null" json:"username"`
    Content string `gorm:"type:text;not null" json:"content"`
    Type    string `gorm:"not null" json:"type"`
}
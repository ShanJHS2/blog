package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Level    int    `gorm:"not null"`
    Avatar   string
}
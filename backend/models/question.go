package models

import "gorm.io/gorm"

type Question struct {
    gorm.Model
    Author  string `gorm:"not null"`
    Content string `gorm:"type:text;not null"`
    Answer  string `gorm:"type:text"`
}
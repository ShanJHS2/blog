package models

import (
	"gorm.io/gorm"
)

type BlogArticle struct {
    gorm.Model    // 自动包含 ID, CreatedAt, UpdatedAt, DeletedAt
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Tags      []string  `json:"tags" gorm:"type:json"`
    Image     string    `json:"image"`
}

type ResearchArticle struct {
    gorm.Model    // 自动包含 ID, CreatedAt, UpdatedAt, DeletedAt
    Title     string    `json:"title"`
    Abstract  string    `json:"abstract"`
    Tags      []string  `json:"tags" gorm:"type:json"`
    Image     string    `json:"image"`
}

type ProjectArticle struct {
    gorm.Model    // 自动包含 ID, CreatedAt, UpdatedAt, DeletedAt
    Title     string    `json:"title"`
    Status    string    `json:"status"`
    Tags      []string  `json:"tags" gorm:"type:json"`
    Image     string    `json:"image"`
}
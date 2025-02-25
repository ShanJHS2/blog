package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/config"
)

// 获取文章列表
func GetArticles(c *gin.Context) {
    page, err := strconv.Atoi(c.Query("page"))
    if err != nil || page < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
        return
    }
    limit, err := strconv.Atoi(c.Query("limit"))
    if err != nil || limit < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
        return
    }
    articleType := c.Query("type")
    var articles interface{}
    offset := (page - 1) * limit

    switch articleType {
    case "blog":
        var blogArticles []models.BlogArticle
        if err := config.DB.Select("id", "title", "tags", "image", "created_at").Order("created_at DESC").Offset(offset).Limit(limit).Find(&blogArticles).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get articles"})
            return
        }
        articles = blogArticles
    case "research":
        var researchArticles []models.ResearchArticle
        if err := config.DB.Select("id", "title", "tags", "image", "created_at").Order("created_at DESC").Offset(offset).Limit(limit).Find(&researchArticles).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get articles"})
            return
        }
        articles = researchArticles
    case "project":
        var projectArticles []models.ProjectArticle
        if err := config.DB.Select("id", "title", "tags", "image", "created_at").Order("created_at DESC").Offset(offset).Limit(limit).Find(&projectArticles).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get articles"})
            return
        }
        articles = projectArticles
    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type parameter"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": articles})
}

// 获取文章数量
func GetArticleCount(c *gin.Context) {
    articleType := c.Query("type")
    var count int64

    switch articleType {
    case "blog":
        if err := config.DB.Model(&models.BlogArticle{}).Count(&count).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get article count"})
            return
        }
    case "research":
        if err := config.DB.Model(&models.ResearchArticle{}).Count(&count).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get article count"})
            return
        }
    case "project":
        if err := config.DB.Model(&models.ProjectArticle{}).Count(&count).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get article count"})
            return
        }
    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type parameter"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"num": count})
}

// 获取文章详情
func GetArticleById(c *gin.Context) {
    id := c.Param("id")
    articleType := c.Query("type")
    var article interface{}

    switch articleType {
    case "blog":
        var blogArticle models.BlogArticle
        if err := config.DB.First(&blogArticle, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
            return
        }
        article = blogArticle
    case "research":
        var researchArticle models.ResearchArticle
        if err := config.DB.First(&researchArticle, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
            return
        }
        article = researchArticle
    case "project":
        var projectArticle models.ProjectArticle
        if err := config.DB.First(&projectArticle, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
            return
        }
        article = projectArticle
    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type parameter"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": article})
}

// 创建文章
func CreateArticle(c *gin.Context) {
    var article interface{}
    articleType := c.Query("type")

    switch articleType {
    case "blog":
        var blogArticle models.BlogArticle
        if err := c.ShouldBindJSON(&blogArticle); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        article = blogArticle
    case "research":
        var researchArticle models.ResearchArticle
        if err := c.ShouldBindJSON(&researchArticle); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        article = researchArticle
    case "project":
        var projectArticle models.ProjectArticle
        if err := c.ShouldBindJSON(&projectArticle); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        article = projectArticle
    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type parameter"})
        return
    }

    if err := config.DB.Create(article).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Article created"})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
    id := c.Param("id")
    articleType := c.Query("type")

    switch articleType {
    case "blog":
        if err := config.DB.Delete(&models.BlogArticle{}, id).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
            return
        }
    case "research":
        if err := config.DB.Delete(&models.ResearchArticle{}, id).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
            return
        }
    case "project":
        if err := config.DB.Delete(&models.ProjectArticle{}, id).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
            return
        }
    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type parameter"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})
}
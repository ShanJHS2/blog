package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// 获取媒体
func GetMedia(c *gin.Context) {
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
    mediaType := c.Query("type")
    if mediaType != "novels" && mediaType != "books" && mediaType != "movies" {
        c.JSON(http.StatusBadRequest, gin.H{"error": mediaType + " is not a valid media type"})
        return
    }

    var media []models.Media
    offset := (page - 1) * limit

    if err := config.DB.Where("type = ?", mediaType).Order("created_at DESC").Offset(offset).Limit(limit).Find(&media).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get media"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": gin.H{
            "medias": media,
        },
    })
}

func CreateMedia(c *gin.Context) {
    var media models.Media
    if err := c.ShouldBindJSON(&media); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&media).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create media"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Media created"})
}

func UpdateMedia(c *gin.Context) {
    var media models.Media
    mediaID := c.Param("mediaId")
    if err := c.ShouldBindJSON(&media); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Model(&models.Media{}).Where("id = ?", mediaID).Updates(media).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update media"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Media updated"})
}

func DeleteMedia(c *gin.Context) {
    mediaID := c.Param("mediaId")
    if err := config.DB.Where("id = ?", mediaID).Delete(&models.Media{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete media"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Media deleted"})
}
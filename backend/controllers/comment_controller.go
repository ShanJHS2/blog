package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/config"
)

func GetComments(c *gin.Context) {
    var comments []models.Comment
    blogID := c.Param("blogID")
    commentType := c.Query("type")

    if err := config.DB.Where("blog_id = ? AND type = ?", blogID, commentType).Find(&comments).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": comments})
}

func CreateComment(c *gin.Context) {
    var comment models.Comment
    if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&comment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Comment created"})
}

// 删除评论
func DeleteComment(c *gin.Context) {
    BlogID := c.Param("blogID")
    commentType := c.Query("type")

    if err := config.DB.Where("blog_id = ? AND type = ?", BlogID, commentType).Delete(&models.Comment{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}
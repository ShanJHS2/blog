package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/config"
	"strconv" 
)

func GetQuestions(c *gin.Context) {
    var questions []models.Question
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

    if err := config.DB.Offset((page - 1) * limit).Limit(limit).Find(&questions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get questions"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"questions": questions})
}

func CreateQuestion(c *gin.Context) {
    var question models.Question
    if err := c.ShouldBindJSON(&question); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&question).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Question created"})
}

func AnswerQuestion(c *gin.Context) {
    var question models.Question
    questionID := c.Param("questionId")
    if err := c.ShouldBindJSON(&question); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Model(&models.Question{}).Where("id = ?", questionID).Update("answer", question.Answer).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to answer question"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Question answered"})
}
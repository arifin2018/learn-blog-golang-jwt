package controllers

import (
	"Blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTags(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tags []models.TagPost
	db.Preload("Post.User").Preload("Post.Image").Preload("Post.Tag").Preload("Post.Comment.User").Find(&tags)

	c.JSON(http.StatusOK, gin.H{"data": tags})
}

func CreateTags(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	var tags models.Tag

	if err := c.ShouldBindJSON(&tags); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	db.Debug().Create(&tags)
	c.JSON(http.StatusOK, gin.H{"data": tags})
}
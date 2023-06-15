package controllers

import (
	"Blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllTags godoc
// @Summary Get all GetTags.
// @Description Get a list of GetTags.
// @Tags Tags
// @Produce json
// @Success 200 {object} []models.TagPost
// @Router /Tags [get]
func GetTags(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tags []models.TagPost
	db.Preload("Post.User").Preload("Post.Image").Preload("Post.Tag").Preload("Post.Comment.User").Find(&tags)

	c.JSON(http.StatusOK, gin.H{"data": tags})
}

// CreateTags godoc
// @Summary Create New Tags.
// @Description Creating a new Tags.
// @Tags Tags
// @Param Body body models.Tag true "the body to create a new Tags"
// @Produce json
// @Success 200 {object} models.Tag
// @Router /Tags [post]
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
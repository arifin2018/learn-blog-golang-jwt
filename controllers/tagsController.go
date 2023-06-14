package controllers

import (
	"Blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
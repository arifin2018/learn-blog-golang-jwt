package controllers

import (
	"Blog/helpers"
	"Blog/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetComments(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Comment []models.GetComment

	db.Preload("User").Find(&Comment)
	c.JSON(http.StatusOK, gin.H{"data": Comment})
}

func CreateComments(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	var Comments models.Comment
	if err := c.ShouldBindJSON(&Comments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	inputComment := models.Comment{Content: Comments.Content, User_id: int(user_id),Created_at: time.Now()}
	db.Debug().Create(&inputComment)

	post_id,_ := strconv.Atoi(c.Param("id"))
	
	inputPostComment := models.PostComment{PostId: post_id, CommentId: inputComment.ID}
	db.Debug().Create(&inputPostComment)

	c.JSON(http.StatusOK, gin.H{"data": inputComment})
}
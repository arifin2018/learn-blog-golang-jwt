package controllers

import (
	"Blog/helpers"
	"Blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePosts(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var Post models.Post

	if err := c.ShouldBindJSON(&Post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	user_id, err := helpers.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	inputPost := models.Post{Content: Post.Content, User_id: int(user_id),Image_id: Post.Image_id,Tag_id: Post.Tag_id, Comment_id: Post.Comment_id}
	db.Create(&inputPost)
	c.JSON(http.StatusOK, gin.H{"data": inputPost})
}

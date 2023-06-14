package controllers

import (
	"Blog/helpers"
	"Blog/models"
	"net/http"
	"time"

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
	inputPost := models.Post{Content: Post.Content, User_id: int(user_id), Created_at: time.Now()}
	db.Debug().Create(&inputPost)
	if len(Post.Image_url) > 0 {
		for _, v := range Post.Image_url {
			ImageCreate := models.Image{Image_url: v}
			db.Debug().Create(&ImageCreate)
			PostImage := models.PostImage{PostId: inputPost.ID,ImageId: ImageCreate.ID}
			db.Debug().Create(&PostImage)
		}
	}
	if len(Post.TagId) > 0 {
		for _, v := range Post.TagId {
			PostTag := models.PostTag{PostId: inputPost.ID, TagId: v}
			db.Debug().Create(&PostTag)
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": inputPost})
}

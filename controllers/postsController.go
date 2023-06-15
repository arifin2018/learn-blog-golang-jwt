package controllers

import (
	"Blog/helpers"
	"Blog/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllPosts godoc
// @Summary Get all GetPosts.
// @Description Get a list of GetPosts.
// @Tags Posts
// @Produce json
// @Success 200 {object} []models.GetPost
// @Router /posts [get]
func GetPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var posts []models.GetPost

	db.Preload("User").Preload("Image").Preload("Tag").Preload("Comment.User").Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// CreatePosts godoc
// @Summary Create New Posts.
// @Description Creating a new Posts.
// @Tags Posts
// @Param Body body models.PostReq true "the body to create a new Posts"
// @Produce json
// @Success 200 {object} models.Post
// @Router /posts [post]
func CreatePosts(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var Post models.Post
	var Tag models.Tag

	if err := c.ShouldBindJSON(&Post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	user_id, err := helpers.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	inputPost := models.Post{Title: Post.Title ,Content: Post.Content, User_id: int(user_id), Created_at: time.Now()}
	
	tx := db.Begin()
	if err := tx.Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
		return
	}
	if err := tx.Debug().Create(&inputPost).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to create input post"})
		return
	}
	
	if len(Post.Image_url) > 0 {
		for _, v := range Post.Image_url {
			ImageCreate := models.Image{Image_url: v}
			if err := tx.Debug().Create(&ImageCreate).Error; err !=nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{"err": err})
				return
			}
			
			PostImage := models.PostImage{PostId: inputPost.ID,ImageId: ImageCreate.ID}
			if err := tx.Debug().Create(&PostImage).Error; err !=nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{"err": err})
				return
			}
		}
	}
	if len(Post.TagId) > 0 {
		for _, v := range Post.TagId {
			PostTag := models.PostTag{PostId: inputPost.ID, TagId: v}
			if err := db.Find(&Tag,v).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{"err": "Tag not found"})
				return
			}
			if err := tx.Debug().Create(&PostTag).Error; err !=nil {
				tx.Rollback()
				c.JSON(http.StatusOK, gin.H{"err": err})
				return
			}
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"data": inputPost})
}

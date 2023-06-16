package controllers

import (
	"Blog/helpers"
	"Blog/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllComments godoc
// @Summary Get all GetComments.
// @Description Get a list of GetComments.
// @Tags Comments
// @Produce json
// @Success 200 {object} []models.GetComment
// @Router /comments/:id [get]
func GetComments(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Comment []models.GetComment

	db.Preload("User").Preload("Image").Find(&Comment)
	c.JSON(http.StatusOK, gin.H{"data": Comment})
}

// CreateComments godoc
// @Summary Create New Comments.
// @Description Creating a new Comments.
// @Tags Comments
// @Param Body body models.Comment true "the body to create a new Comment"
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comments/:id [post]
func CreateComments(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	tx := db.Begin()

	if err := tx.Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
		return
	}

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
	if err := tx.Debug().Create(&inputComment).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to create input comment"})
		return
	}

	post_id,_ := strconv.Atoi(c.Param("id"))
	
	inputPostComment := models.PostComment{PostId: post_id, CommentId: inputComment.ID}
	if err := tx.Debug().Create(&inputPostComment).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to create input Post Comment"})
		return
	}

	for _, v := range Comments.ImageUrl {
		inputImages := models.Image{Image_url: v}
		if err := tx.Debug().Create(&inputImages).Error; err != nil {
			tx.Rollback()
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": "failed to create input Image"})
			return
		}
		inputCommentImage := models.CommentImage{CommentId: inputComment.ID, ImageId:inputImages.ID}
		if err := tx.Debug().Create(&inputCommentImage).Error; err != nil {
			tx.Rollback()
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"err": "failed to create input Comment Image"})
			return
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"data": inputComment})
}

// DeleteComments godoc
// @Summary Delete a Comments.
// @Description Delete a Comments.
// @Tags Comments
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comments/:id [delete]
func DeleteComments(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	db.Delete(&models.Comment{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": []models.Comment{}})
}
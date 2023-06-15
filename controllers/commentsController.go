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

// GetAllComments godoc
// @Summary Get all GetComments.
// @Description Get a list of GetComments.
// @Tags Comments
// @Produce json
// @Success 200 {object} []models.GetComment
// @Router /Comments/:id [get]
func GetComments(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Comment []models.GetComment

	db.Preload("User").Find(&Comment)
	c.JSON(http.StatusOK, gin.H{"data": Comment})
}

// CreateComments godoc
// @Summary Create New Comments.
// @Description Creating a new Comments.
// @Tags Comments
// @Param Body body models.Comment true "the body to create a new Comment"
// @Produce json
// @Success 200 {object} models.Comment
// @Router /Comments/:id [post]
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
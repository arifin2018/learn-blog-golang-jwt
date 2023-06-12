package controllers

import (
	"Blog/helpers"
	"Blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers()  {
	
}

func RegisterUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var inputUser models.User
	

	if err := c.ShouldBindJSON(&inputUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	user := models.User{}


	if !helpers.ValidEmail(inputUser.Email){
		c.JSON(http.StatusBadRequest, gin.H{"error": "email invalid"})
        return
	}

	user.Email = inputUser.Email
	user.Nickname = inputUser.Nickname
	user.Password = inputUser.Password

	_, err := user.SaveUser(db)

	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	userMapping := map[string]string{
        "Nickname"	: user.Nickname,
        "Email"		: user.Email,
    }

	c.JSON(200, gin.H{
		"message":"user data",
		"user": userMapping,
	})
}
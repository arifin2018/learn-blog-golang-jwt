package controllers

import (
	"Blog/helpers"
	"Blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginUsers(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	var LoginUser models.LoginUsers

	if err := c.ShouldBindJSON(&LoginUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	
	user := models.User{}
	user.Email = LoginUser.Email
	user.Password = LoginUser.Password
	// result := map[string]interface{}{}
	// db.Model(&user).Where("email = ?","arifin@lenna.ai").Take(&result)
	token,_,err := user.LoginCheck(user.Email,user.Password,db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	db.Model(&user).Where("email = ?",LoginUser.Email).Update("token",token)
	
	result := map[string]interface{}{}
	db.Model(&user).Where("email = ?",LoginUser.Email).Take(&result)

	delete(result, "id")
	delete(result, "password")

	c.JSON(http.StatusOK, gin.H{"message": "login success", "user": result})
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
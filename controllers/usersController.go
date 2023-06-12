package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetUsers()  {
	
}

func PostUsers(c *gin.Context) {
	// db := c.MustGet("db").(*gorm.DB)

	c.JSON(200, gin.H{
		"message":"ini user",
	})
}
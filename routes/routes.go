package routes

import (
	"Blog/controllers"
	"Blog/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
    })
	
	r.GET("/ping", func(c *gin.Context) {
		helpers.ValidEmail("arifin@gmail.com")
	})

	r.POST("/register", controllers.RegisterUsers)

	return r
}
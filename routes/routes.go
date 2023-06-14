package routes

import (
	"Blog/controllers"
	"Blog/helpers"
	"Blog/middlewares"

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
	r.POST("/login", controllers.LoginUsers)

	MiddlewarePrefixGroupPosts := r.Group("/posts")
	MiddlewarePrefixGroupPosts.Use(middlewares.JwtAuthMiddleware())
	MiddlewarePrefixGroupPosts.POST("/", controllers.CreatePosts)

	MiddlewarePrefixGroupTags := r.Group("/Tags")
	MiddlewarePrefixGroupTags.Use(middlewares.JwtAuthMiddleware())
	MiddlewarePrefixGroupTags.POST("/", controllers.CreateTags)


	return r
}
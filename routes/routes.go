package routes

import (
	"Blog/controllers"
	"Blog/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
    })
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"name": "arifin"})
	})
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/register", controllers.RegisterUsers)
	r.POST("/login", controllers.LoginUsers)

	MiddlewarePrefixGroupPosts := r.Group("/posts")
	MiddlewarePrefixGroupPosts.Use(middlewares.JwtAuthMiddleware())
	MiddlewarePrefixGroupPosts.GET("/", controllers.GetPosts)
	MiddlewarePrefixGroupPosts.POST("/", controllers.CreatePosts)

	MiddlewarePrefixGroupTags := r.Group("/tags")
	MiddlewarePrefixGroupTags.Use(middlewares.JwtAuthMiddleware())
	MiddlewarePrefixGroupTags.GET("/", controllers.GetTags)
	MiddlewarePrefixGroupTags.POST("/", controllers.CreateTags)
	MiddlewarePrefixGroupTags.PUT("/:id", controllers.UpdateTags)
	MiddlewarePrefixGroupTags.DELETE("/:id", controllers.DeleteTags)

	MiddlewarePrefixGroupComments := r.Group("/comments")
	MiddlewarePrefixGroupComments.Use(middlewares.JwtAuthMiddleware())
	MiddlewarePrefixGroupComments.GET("/:id", controllers.GetComments)
	MiddlewarePrefixGroupComments.POST("/:id", controllers.CreateComments)
	MiddlewarePrefixGroupComments.DELETE("/:id", controllers.DeleteComments)

	return r
}
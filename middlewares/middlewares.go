package middlewares

import (
	"Blog/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.TokenValid(c)
		if err != nil {
			log.Println("it was something wrong")
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
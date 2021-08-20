package main

import (
	"app/handler"
	"app/middleware"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		names := []string{"John", "Jane"}
		c.SecureJSON(http.StatusOK, names)
	})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.POST("/login", handler.Login)

	authGroup := router.Group("/auth")
	authGroup.Use(middleware.LoginCheck())
	{
		authGroup.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!!",
			})
		})
	}

	router.Run("127.0.0.1:8080")
}

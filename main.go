package main

import (
	"app/handler"
	"app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome!!",
		})
	})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.POST("/login", handler.Login)
	router.POST("/registration", handler.Registration)

	authGroup := router.Group("/auth")
	authGroup.Use(middleware.LoginCheck())
	{
		authGroup.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!!",
			})
		})

		authGroup.POST("/logout", handler.Logout)
	}

	router.Run("127.0.0.1:8080")
}

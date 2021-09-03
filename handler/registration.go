package handler

import (
	"app/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Registration(c *gin.Context) {
	uid := c.PostForm("uid")
	password := c.PostForm("password")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	userSchema := model.User{Uid: uid, Password: string(hash)}
	e := userSchema.RegisterUser()
	if e != nil {
		c.JSON(500, gin.H{
			"message": e.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "User register was succeed.",
	})
}

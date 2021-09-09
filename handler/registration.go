package handler

import (
	"app/aop"
	"app/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	db, err := aop.Connect()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	e := aop.ExecuteCreateQuery(db, func(tx *gorm.DB) error {
		return userSchema.CreateUser(tx)
	})

	if e != nil {
		c.JSON(500, gin.H{
			"message": e.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "User register was succeed.",
		})
	}
}

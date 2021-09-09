package handler

import (
	"app/aop"
	"app/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var requestBody model.UserRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 3)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	userSchema := model.User{Uid: requestBody.Uid, Password: string(hash)}
	db, err := aop.Connect()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	e := aop.ExecuteCreateQuery(db, func(tx *gorm.DB) error {
		return userSchema.CreateUser(tx)
	})

	if e != nil {
		c.JSON(500, gin.H{
			"error": e.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "User register was succeed.",
		})
		return
	}
}

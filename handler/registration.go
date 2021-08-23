package handler

import (
	"app/aop"
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Registration(c *gin.Context) {
	uid := c.PostForm("uid")
	password := c.PostForm("password")

	db, err := aop.Connect()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	user := model.User{Uid: uid, Password: string(hash)}

	db.Create(&user)
	c.Status(http.StatusOK)
}

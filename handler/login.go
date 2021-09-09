package handler

import (
	"app/model"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	/*
		uid: string ユーザーID
		password: string パスワード
	*/
	var requestBody model.UserRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userSchema := model.User{Uid: requestBody.Uid}
	if err := userSchema.GetUserById(); err != nil {
		c.JSON(403, gin.H{
			"message": err.Error(),
		})
		return
	}

	//パスワードを比較
	if err := bcrypt.CompareHashAndPassword([]byte(userSchema.Password), []byte(requestBody.Password)); err != nil {
		log.Println(err)
		c.JSON(403, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		//sessionを取得
		var session = sessions.Default(c)

		var option = sessions.Options{SameSite: http.SameSiteLaxMode}

		session.Options(option)

		//session idとuidを紐付け
		session.Set("sessionId", userSchema.Uid)
		session.Save()
	}
	c.JSON(200, gin.H{
		"message": "Your login was succeed.",
	})
}

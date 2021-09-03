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
		uid: string ユーザーID Repository実装までは使用しない
		password: string パスワード
	*/
	var uid = c.PostForm("uid")
	var password = c.PostForm("password")

	userSchema := model.User{Uid: uid}
	if err := userSchema.GetUserById(); err != nil {
		log.Println(err)
		c.JSON(403, gin.H{
			"message": err.Error(),
		})
	}

	//パスワードを比較
	if err := bcrypt.CompareHashAndPassword([]byte(userSchema.Password), []byte(password)); err != nil {
		log.Fatalln(err)
		c.Status(http.StatusForbidden)
	} else {
		//sessionを取得
		var session = sessions.Default(c)

		var option = sessions.Options{SameSite: http.SameSiteLaxMode}

		session.Options(option)

		//session idとuidを紐付け
		session.Set("sessionId", userSchema.Uid)
		session.Save()
	}
}

package controller

import (
	"app/model"
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
	var _ = c.PostForm("uid")
	var password = c.PostForm("password")

	//仮のRepository
	user := model.User{Uid: "user1", Password: "0000"}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	//パスワードを比較
	err = bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		c.Status(http.StatusForbidden)
	} else {
		//sessionを取得
		var session = sessions.Default(c)

		//session idとuidを紐付け
		session.Set("sessionId", user.Uid)
		session.Save()
	}
}

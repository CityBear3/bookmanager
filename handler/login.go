package handler

import (
	"app/aop"
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

	db, err := aop.Connect()
	if err != nil {
		c.Status(http.StatusUnauthorized)
	}

	dbCloser, err := db.DB()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	defer dbCloser.Close()

	user := model.User{Uid: uid}
	db.Find(&user)
	db.Logger.LogMode(4)

	//パスワードを比較
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusForbidden)
	} else {
		//sessionを取得
		var session = sessions.Default(c)

		var option = sessions.Options{SameSite: http.SameSiteLaxMode}

		session.Options(option)

		//session idとuidを紐付け
		session.Set("sessionId", user.Uid)
		session.Save()
	}
}

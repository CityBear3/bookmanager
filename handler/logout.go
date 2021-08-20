package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	//sessionを取得
	session := sessions.Default(c)

	//sessionをクリア
	session.Clear()
	session.Save()

	c.Status(http.StatusOK)
}

package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	//sessionを取得
	session := sessions.Default(c)

	//sessionをクリア
	session.Clear()
	err := session.Save()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "Your logout was succeed.",
	})
}

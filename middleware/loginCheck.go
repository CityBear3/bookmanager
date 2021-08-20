package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//sessionを取得
		session := sessions.Default(c)

		//sessionからuidを取得
		uid := session.Get("sessionId")

		if uid == nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

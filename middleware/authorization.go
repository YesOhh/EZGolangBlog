package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goBlog/mylog"
	"net/http"
	"strings"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/login") || strings.HasPrefix(c.Request.URL.Path, "/register") {
			c.Next()
		}else if strings.HasPrefix(c.Request.URL.Path, "/static") {
			c.Next()
		} else {
			session := sessions.Default(c)
			authenticated := session.Get("status")
			mylog.MyLogger.Println("current authenticated: ", authenticated)
			if authenticated == nil || authenticated == false {
				c.Redirect(http.StatusTemporaryRedirect, "/login")
			} else {
				c.Next()
			}
		}
	}
}

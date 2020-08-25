package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 用于恢复错误
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("发生错误： %s\n", err)
				c.HTML(http.StatusInternalServerError, "error.html", gin.H{
					"title": "发生错误",
					"error": "服务器内部发生错误",
				})
			}
		}()
		c.Next()
	}
}

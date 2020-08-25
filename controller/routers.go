package controller

import (
	"github.com/gin-gonic/gin"
	"goBlog/middleware"
	"net/http"
)

func LoadRouters(r *gin.Engine)  {
	// 第二个是存放静态资源的地址，绑到第一个参数上，用于模板中调用
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.Use(middleware.Recovery())

	r.GET("/posts", Articles)
	r.POST("/posts", Articles)
	r.GET("/", Total)
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error.tmpl", gin.H{
			"title": "发生错误",
			"error": "未找到该页面",
		})
	})
}

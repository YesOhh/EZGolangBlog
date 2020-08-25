package main

import (
	"github.com/gin-gonic/gin"
	"goBlog/controller"
	"goBlog/initialization"
)

func main() {
	r := gin.Default()
	controller.LoadRouters(r)
	ip := initialization.Conf.Setting.Ip
	port := initialization.Conf.Setting.Port
	if ip == "" {
		ip = "127.0.0.1"
	}
	if port == "" {
		port = "8080"
	}
	r.Run(ip + ":" + port)
}
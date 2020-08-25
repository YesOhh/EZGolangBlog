package main

import (
	"github.com/gin-gonic/gin"
	"goBlog/controller"
)

func main() {
	r := gin.Default()
	controller.LoadRouters(r)
	r.Run()
}
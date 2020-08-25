package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goBlog/model"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	log.Println("method:", c.Request.Method, "url:", c.Request.URL.Path)
	var user model.UserModel
	if err := c.ShouldBind(&user); err != nil {
		log.Println("参数错误，", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"result": "failure",
			"message": "邮箱格式错误",
		})
		return
	}

	if user.ValidUser() {
		session := sessions.Default(c)
		session.Set("status", true)
		session.Save()
		c.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "failure",
			"message": "用户名不存在或密码错误",
		})
	}
}

func Register(c *gin.Context) {
	log.Println("method:", c.Request.Method, "url:", c.Request.URL.Path)
	var registerUser model.RegisterModel
	if err := c.ShouldBind(&registerUser); err != nil {
		log.Println("参数错误，", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"result": "failure",
			"message": "邮箱格式错误",
		})
		return
	}
	if registerUser.Password != registerUser.PasswordAgain {
		c.JSON(http.StatusOK, gin.H{
			"result": "failure",
			"message": "两次输入密码不相同",
		})
		return
	}
	user := model.UserModel{Password: registerUser.Password, Email: registerUser.Email, Nickname: registerUser.Nickname}
	queryUser := user.QueryUserByEmail()
	if queryUser.Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "failure",
			"message": "该邮箱已注册",
		})
	} else {
		user.SaveUser()
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

func LogOut(c *gin.Context) {
	log.Println("method:", c.Request.Method, "url:", c.Request.URL.Path)
	session := sessions.Default(c)
	if session.Get("status") == true {
		session.Set("status", false)
		session.Save()
		log.Println("authenticated: ", session.Get("status"))
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}
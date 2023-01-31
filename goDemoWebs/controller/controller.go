package controller

import (
	"fmt"
	"golang/goDemoWebs/dao"
	"golang/goDemoWebs/model"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	uname := c.PostForm("username")
	pword := c.PostForm("password")

	user := model.User{
		Username: uname,
		Password: pword,
	}

	dao.Mgr.RegisterUser(&user)

	c.Redirect(301, "/")
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(200, "userlist.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)

	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在")
		fmt.Println("用户名不存在")
	} else {
		if u.Password != password {
			c.HTML(200, "login.html", "密码错误")
			fmt.Println("密码错误")
		} else {
			c.Redirect(301, "/")
			fmt.Println("登陆成功")
		}
	}
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

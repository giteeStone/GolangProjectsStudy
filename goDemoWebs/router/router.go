package router

import (
	"golang/goDemoWebs/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()

	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.POST("/register", controller.RegisterUser)
	e.GET("/register", controller.GoRegister)
	e.GET("/", controller.Index)

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	e.Run()
}

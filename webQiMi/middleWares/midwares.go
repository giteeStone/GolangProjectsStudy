/*Gin中的中间件必须是一个gin.HandlerFunc类型。
Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。
这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，
比如登录认证、权限校验、数据分页、记录日志、耗时统计等。*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index 进来了")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
	c.Next() //c.Abort()  Next()是handler后面所有的！！
}

func m1(c *gin.Context) {
	fmt.Println("mid1 进来了")
	c.Next()
}

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

//认证中间件

func authMidware(need bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mid2 in....")
		//判断是否是登录用户
		if need {
			if "是登录用户" == "111" {
				c.Next()
			} else {
				c.Abort()
			}
		}
	}
}

func main() {
	r := gin.Default()

	//注册路由 r.Use(statCost())

	r.GET("/index", m1, indexHandler) // handlers ...HandlerFunc!!!!!!
	r.Run()
}

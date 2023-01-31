//https://www.liwenzhou.com/posts/Go/gin/

// restFul
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get !",
	})

}

func main() {
	r := gin.Default()        //默认路由引擎
	r.GET("/hello", sayHello) //GET访问时，执行sayHello
	r.POST("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Post",
		})
	})

	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "PUT",
		})
	})

	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "DELETE",
		})
	})

	//启动服务
	r.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})

	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(200, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
			//....
		}
	})

	//默认无地址
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})

	//路由组
	/*func main() {
		r := gin.Default()
		userGroup := r.Group("/user")
		{
			userGroup.GET("/index", func(c *gin.Context) {...})
			userGroup.GET("/login", func(c *gin.Context) {...})
			userGroup.POST("/login", func(c *gin.Context) {...})

		}
		shopGroup := r.Group("/shop")
		{
			shopGroup.GET("/index", func(c *gin.Context) {...})
			shopGroup.GET("/cart", func(c *gin.Context) {...})
			shopGroup.POST("/checkout", func(c *gin.Context) {...})
		}
		r.Run()



		路由组也是支持嵌套的，例如：

	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		shopGroup.GET("/cart", func(c *gin.Context) {...})
		shopGroup.POST("/checkout", func(c *gin.Context) {...})
		// 嵌套路由组
		xx := shopGroup.Group("xx")
		xx.GET("/oo", func(c *gin.Context) {...})
	}
	}*/

	r.Run()
}

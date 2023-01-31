package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	r := gin.Default()
	r.Use(favicon.New("./"))
	r.GET("/json", func(c *gin.Context) {
		// data := map[string]interface{}{
		// 	"name":    "小王子",
		// 	"message": "helloworld",
		// 	"age":     18,
		// }

		data1 := gin.H{
			"name":    "小王子",
			"message": "helloworld",
			"age":     18,
		}
		c.JSON(200, data1)
	})
	r.Run()
}

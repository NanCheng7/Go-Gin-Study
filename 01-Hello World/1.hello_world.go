package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个默认的路由

	router := gin.Default()

	router.GET("/index", func(c *gin.Context) {
		c.String(200, "hello NanCheng_7")
	})

	router.Run("0.0.0.0:8080")
}

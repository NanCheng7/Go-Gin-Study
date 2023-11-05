package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.Default()
	//请求头的各种获取方式
	router.GET("/", func(c *gin.Context) {
		// 字母大小写不区分 单词与单词之间是用 - 连接
		// 获取单个请求头
		fmt.Println(c.GetHeader("User-Agent"))
		fmt.Println(c.GetHeader("user-agent"))
		fmt.Println(c.GetHeader("user-Agent"))

		//获取全部请求头
		fmt.Println(c.Request.Header)
		fmt.Println(c.Request.Header["User-Agent"])
		// 没有通过 get方法进行转换，导致无法无视大小写
		fmt.Println(c.Request.Header["User-agent"])

		//如果是使用的自定义的请求头，也是可以使用get方法无视大小写
		fmt.Println(c.GetHeader("token"))
		fmt.Println(c.GetHeader("Token"))
		c.JSON(200, gin.H{"msg": "成功"})
	})

	//爬虫和用户区别对待
	router.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")

		if strings.Contains(userAgent, "python") {
			// 爬虫
			c.JSON(0, gin.H{"msg": "爬虫"})
		}
		c.JSON(0, gin.H{"msg": "用户"})

	})

	router.Run(":80")
}

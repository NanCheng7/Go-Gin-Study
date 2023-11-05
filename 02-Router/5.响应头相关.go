package main

import "github.com/gin-gonic/gin"

// _setRespHeader 设置响应头
func _setRespHeader(c *gin.Context) {
	c.Header("token", "NanCheng")
	// 在这里修改content-type可以修改响应的数据类型，这里直接放回text，浏览器会把他当作文本下载
	c.Header("content-type", "application/text; charset=utf-8")
	c.JSON(0, gin.H{"msg": "看看响应头"})
}

func main() {
	router := gin.Default()

	router.GET("/resp", _setRespHeader) // 设置响应头

	router.Run(":80")
}

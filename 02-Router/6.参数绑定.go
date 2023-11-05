package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfo6 struct {
	Name string `json:"name" form:"name" xml:"name"`
	Age  int    `json:"age" form:"age" xml:"age"`
	Sex  string `json:"sex" form:"sex" xml:"sex"`
}

// ShouldBind JSON
func _ShouldBindJSON(c *gin.Context) {
	var userInfo UserInfo6
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": "你错了"})
		return
	}
	c.JSON(200, userInfo)
}

// ShouldBind Query
func _ShouldBindQuery(c *gin.Context) {
	var userInfo UserInfo6
	err := c.ShouldBindQuery(&userInfo)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"msg": "你错了"})
		return
	}
	c.JSON(200, userInfo)
}

// ShouldBind Uri
func _ShouldBindUri(c *gin.Context) {
	var userInfo UserInfo6
	err := c.ShouldBindUri(&userInfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": "你错了"})
		return
	}
	c.JSON(200, userInfo)
}

// ShouldBind
func _ShouldBind(c *gin.Context) {
	var userInfo UserInfo6
	err := c.ShouldBind(&userInfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": "你错了"})
	}
	c.JSON(200, userInfo)
}

func main() {
	router := gin.Default()

	//Should Bind JSON
	router.POST("/json", _ShouldBindJSON)

	//Should Bind Query
	router.POST("/query", _ShouldBindQuery)

	//Should Bind Uri
	router.POST("/uri/:name/:age/:sex", _ShouldBindUri)

	//Should Bind
	router.POST("/", _ShouldBind)
	router.Run(":80")
}

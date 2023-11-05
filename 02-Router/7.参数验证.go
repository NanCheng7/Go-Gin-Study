package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Normal 常见的bind验证器
type Normal struct {
	Name       string `json:"name" form:"name" binding:"required,min=4,max=6,len=5"`
	Age        int    `json:"age" form:"age" binding:"gt=18,lt=30"       `
	Password   string `json:"pwd" form:"pwd" binding:"required"`
	RePassword string `json:"repwd" form:"repwd" binding:"required,eqfield=Password"     `
}

type Inner struct {
	Name     string   `json:"name" form:"name" binding:"contains=nancheng,excludes=s,startswith=k,endswith=g"`
	Age      int      `json:"age" form:"age"`
	Sex      string   `json:"sex" form:"sex" binding:"required,oneof=man woman"`
	LikeList []string `json:"like-list" form:"like-list" binding:"required,dive,startswith=like"`
	IP       string   `json:"ip" form:"ip" binding:"ip"`
	Url      string   `json:"url" form:"url" binding:"url"`
	Uri      string   `json:"uri" form:"uri" binding:"uri"`
	Date     string   `json:"date" form:"date" binding:"datetime=2006-01-02 15:04:05"` // 2006年1月2日下午3点4分5秒
}

func _normalBind(c *gin.Context) {
	var userInfo Normal
	err := c.ShouldBind(&userInfo)

	fmt.Println(userInfo)
	if err != nil {
		c.JSON(200, err.Error())
		fmt.Println(err.Error())
		return
	}
	c.JSON(200, gin.H{"msg": userInfo})
}

func _innerBind(c *gin.Context) {
	var userInfo Inner
	err := c.ShouldBind(&userInfo)

	fmt.Println(userInfo)
	if err != nil {
		c.JSON(200, err.Error())
		fmt.Println(err.Error())
		return
	}
	c.JSON(200, gin.H{"msg": userInfo})

}

func main() {
	router := gin.Default()
	//常见验证器
	router.POST("/normal", _normalBind)
	router.POST("/inner", _innerBind)
	router.Run(":80")
}

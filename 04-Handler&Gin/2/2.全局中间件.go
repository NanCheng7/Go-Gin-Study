package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// #region init func

func __init1(c *gin.Context) {
	// 初始化
	fmt.Println("init1  ..in")
	c.JSON(200, gin.H{"msg": "init success"})
	c.Abort()
	c.Next()
	fmt.Println("init1  ..out")
}
func __init2(c *gin.Context) {
	// 初始化
	fmt.Println("init2  ..in")
	c.JSON(200, gin.H{"msg": "init success"})
	c.Next()
	fmt.Println("init2  ..out")
}

//#endregion

//#region type struct

type UserData struct {
	Name string
	Age  int
}

//#endregion

func _transferData(c *gin.Context) {
	fmt.Println("User upload a data")
	data := UserData{
		"NanCheng",
		20,
	}
	c.Set("userData", data)
}
func main() {

	router := gin.Default()

	//router.Use(__init1, __init2)
	router.GET("/1", func(c *gin.Context) {
		fmt.Println(1)
		c.JSON(200, gin.H{"11": "11"})
	})
	router.GET("/2", func(c *gin.Context) {
		fmt.Println(2)
		c.JSON(200, gin.H{"22": "2"})
	})
	// tD ==> TransferData
	router.GET("/tD", _transferData, func(c *gin.Context) {
		//未断言类型
		tUserdata, _ := c.Get("userData")
		c.JSON(200, gin.H{"data": tUserdata})
		// 如果我想要使用数据结构体里面的一个字段，就必须要对它进行断言操作才可访问
		// 断言类型 ==> userData 是否断言成功 ==> is
		userData, is := tUserdata.(UserData)
		if is {
			c.JSON(200, gin.H{"data": userData.Name})
		}

	})

	router.Run(":80")

}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//#region sad

//#endregion

func _singleRouterHandler1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World1",
	})
	c.Abort()
	fmt.Println("aa: 11")
	fmt.Println("aa: 22")

}
func _singleRouterHandler2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World2",
	})
}
func _singleRouterHandler3(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World3",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", _singleRouterHandler1, _singleRouterHandler2, _singleRouterHandler3)

	router.Run(":8080")
}

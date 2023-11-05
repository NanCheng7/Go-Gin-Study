package main

import "github.com/gin-gonic/gin"

func _singleRouterHandler1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World1",
	})
	c.Abort()

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

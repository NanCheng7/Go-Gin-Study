package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserName string `json:"user_name"`
	Age int `json:"age"`
	Password string `json:"-"`	//忽略转换为JSON
}


func _string(c *gin.Context) {
	c.String(http.StatusOK, "响应你好")
}
func _json(c *gin.Context) {

	//user := UserInfo{
	//	"NanCheng",
	//	16,
	//	"122345",
	//}
	//c.JSON(200,user)

	// JSON 响应map
	//userMap := map[string]string{
	//	"UserName":"NanCheng",
	//	"age":"23",
	//}
	//
	//c.JSON(http.StatusOK,userMap)

	// 直接响应json
	c.JSON(200,gin.H{"user_name":"NanCheng","age":12})

}

func _jsonP(c *gin.Context) {
	a := &UserInfo{
		"NanCheng",
		16,
		"111222",
	}
	c.JSONP(200,a)
}

func main() {
	router := gin.Default()
	router.GET("/", _string)
	router.GET("/json", _json)
	router.GET("/jsonP",_jsonP)
	router.Run(":80")
}

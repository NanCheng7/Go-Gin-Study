package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Password string `json:"-"` //忽略转换为JSON
}

func _string(c *gin.Context) {
	c.String(http.StatusOK, "响应你好")
}
func _json(c *gin.Context) {

	//user := UserInfo_8{
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
	c.JSON(200, gin.H{"user_name": "NanCheng", "age": 12})

}

func _jsonP(c *gin.Context) {
	a := &UserInfo{
		"NanCheng",
		16,
		"111222",
	}
	c.JSONP(200, a)
}

func _xml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"user": "NanCheng", "message": "this is key", "Status": http.StatusOK})
}

func _yml(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"user": "NanCheng", "Message": "this is yml"})
}

func _html(c *gin.Context) {
	//
	//type User struct {
	//	username string `json:"username"`
	//	//message  string `json:"message"`
	//	//password string `json:"-"`
	//}
	//user := User{
	//	"NanCheng",
	//	//"studying go",
	//	//"123456",
	//}

	c.HTML(200, "HelloRouter.html", gin.H{"username": "NanCheng"})
	//c.HTML(200, "HelloRouter.html", user)
}

// 重定向
func _redirect(c *gin.Context) {
	c.Redirect(302, "http://www.baidu.com")
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("HTML-Templates/router/*")
	// 在golang中，没有相对文件的路径，只有相对项目的路径

	// 配置单个文件，网页请求的路由，文件的路径
	router.StaticFile("/static/people", "./static/People.png")
	// 网页请求这个静态目录的前缀，第二个参数是一个目录，注意，前缀不可以重复
	router.StaticFS("/static/sss", http.Dir("./static/static"))

	router.GET("/", _string)
	router.GET("/json", _json)
	router.GET("/jsonP", _jsonP)
	router.GET("/xml", _xml)
	router.GET("/yml", _yml)
	router.GET("/html", _html)
	router.GET("/baidu", _redirect)
	router.Run(":80")
}

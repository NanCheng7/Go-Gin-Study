package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _userListView(c *gin.Context) {
	var userList []UserInfo = []UserInfo{
		{"NanCheng", 666}, {"BeiXiang", 14},
	}
	c.JSON(200, Response{0, userList, "success"})

}

func _lgHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "john",
		"age":  20,
	})
}

// #region 路由初始化
func _routerInit(router *gin.Engine) {
	api := router.Group("api")
	_userRouterInit(api)
	_loginRouterInit(api)
}

//#endregion

// #region 用户路由初始化
func _userRouterInit(router *gin.RouterGroup) {
	userController := router.Group("userController")
	{
		userController.GET("/home", _lgHome)
		userController.GET("/userList", _userListView)
	}
}

//#endregion

// #region 登录路由初始化
func _loginRouterInit(router *gin.RouterGroup) {
	// 登录接口
	loginController := router.Group("loginController")
	{
		loginController.GET("/api", _lgHome)
	}
}

//#endregion

func main() {

	router := gin.Default()
	_routerInit(router)
	router.Run(":8080")
}

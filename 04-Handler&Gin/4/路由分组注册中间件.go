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

// #region 路由初始化
func _routerInit(router *gin.Engine) {
	api := router.Group("api")
	_userRouterInit(api)
	_loginRouterInit(api)
}

//#endregion

// #region 登陆路由初始化
func _loginRouterInit(router *gin.RouterGroup) {
	loginController := router.Group("loginController").Use(_Middleware("登录验证失败"))
	{
		loginController.GET("/login", func(c *gin.Context) {
			c.JSON(200, Response{0, nil, "success"})
		})
	}
}

//#endregion

//#region 用户路由初始化

func _userRouterInit(router *gin.RouterGroup) {
	userController := router.Group("userController").Use(_Middleware("用户登陆失败"))
	{
		userController.GET("/userList", _userListView)
	}
}

//#endregion

// #region 权限校验中间件
func _Middleware(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		//如果是则校验成功
		if token == "1234" {
			c.Next()
			return
		}
		c.JSON(200, Response{-1, nil, msg})
		c.Abort()
	}
}

// #endregion
func main() {

	router := gin.Default()
	_routerInit(router)
	router.Run(":8080")
}

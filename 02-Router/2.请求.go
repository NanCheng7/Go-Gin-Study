package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func _query(c *gin.Context) {
	user := c.Query("user")
	fmt.Println(user)
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.GetQueryArray("user"))
	fmt.Println(c.GetQueryMap("us"))

}

func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))               //访问第一个
	fmt.Println(c.PostFormArray("name"))          //访问列表
	fmt.Println(c.DefaultPostForm("addr", "四川省")) // 如果用户没传，就使用默认值
	forms, err := c.MultipartForm()               // 接收所有的form参数，包括文件
	fmt.Println(forms, err)
}

func _raw(c *gin.Context) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":

		// json解析到结构体
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		err := json.Unmarshal(body, &user)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(user)
	}
}

func bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func main() {
	router := gin.Default()

	router.GET("/query", _query)
	router.GET("/param/:user_id", _param)
	router.GET("/param/:user_id/:book_id", _param)
	router.GET("/form", _form)
	router.GET("/row", _raw)
	router.Run(":80")

}

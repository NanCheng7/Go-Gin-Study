package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type UserInfo_8 struct {
	Name string `json:"name" binding:"required" msg:"用户名校验失败"`
	Age  int    `json:"age" binding:"required" msg:"请输入年龄"`
}

// 返回错误信息
func _returnErrorMessage(err error, obj any) string {
	getObj := reflect.TypeOf(obj)
	if errs, isError := err.(validator.ValidationErrors); isError {
		// 断言成功
		for _, e := range errs {
			//循环判断是否有字段报错
			// 如果有久通过反射获取Tag 并且返回Tag的msg
			if field, exist := getObj.Elem().FieldByName(e.Field()); exist {
				msg := field.Tag.Get("msg")
				return msg
			}
		}
	}
	return ""
}

func _customErrorMessage(c *gin.Context) {
	var userInfo UserInfo_8
	err := c.ShouldBindJSON(&userInfo)
	fmt.Println(userInfo)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	// 业务逻辑
	c.JSON(200, gin.H{"data": userInfo})
}

func main() {
	router := gin.Default()

	router.POST("/", _customErrorMessage)

	router.Run(":80")
}

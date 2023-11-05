package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type User9 struct {
	Name string `json:"name" binding:"required,sign" msg:"用户名校验失败"`
	Age  int    `json:"age" binding:"required" msg:"请输入年龄"`
}

func _login(c *gin.Context) {
	var user User9
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(200, gin.H{
			"msg": _returnErrorMessage(err, &user),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": user,
	})
}

func signValid(fl validator.FieldLevel) bool {
	var nameList = []string{"NanCheng", "NanCY"}
	for _, nameStr := range nameList {
		fieldName := fl.Field().Interface().(string)
		if nameStr == fieldName {
			return false
		}
	}
	return true
}
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
func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}

	router.POST("/", _login)

	router.Run(":80")
}

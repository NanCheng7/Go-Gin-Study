package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// _fileLoad 单文件上传
func _fileLoad(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{"code": 200, "msg": "上传文件失败"})
		return
	}
	//第一种保存文件
	//c.SaveUploadedFile(file, "./loadFile/"+file.Filename)

	//第二种
	//获取字节流
	readFile, err := file.Open()
	if err != nil {
		c.JSON(200, gin.H{"code": 200, "msg": "上传文件失败"})
		return
	}
	// 创建文件
	out, err := os.Create("./loadFile/xx.txt")
	// 异常捕获
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)
	// 复制文件流
	_, es := io.Copy(out, readFile)
	if es != nil {
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": "上传文件成功"})

}

// _multiFileLoad 多文件上传
func _multiFileLoad(c *gin.Context) {
	//获取所有的文件 form 实际是一个map[string][]*FileHeader
	form, _ := c.MultipartForm()

	files := form.File["upload[]"]
	for _, file := range files {
		c.SaveUploadedFile(file, "./loadFile/"+file.Filename)
	}

	c.JSON(200, gin.H{"msg": "上传成功! 上传了" + fmt.Sprintf(" %d 个文件", len(files))})
}

func main() {
	router := gin.Default()

	router.POST("/upload", _fileLoad)

	router.POST("/uploads", _multiFileLoad)

	router.Run(":8080")
}

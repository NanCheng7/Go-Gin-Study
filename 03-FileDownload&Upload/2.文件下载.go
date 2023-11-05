package main

import "github.com/gin-gonic/gin"

func _fileDownload(c *gin.Context) {
	c.File("F:\\GO\\Study\\downloadFile\\test.png")
	// 文档形式为8位字节流
	c.Header("Content-Type", "application/octet-stream")
	//内容处置的形式为 ==> 附件形式，文件名为test.png
	c.Header("Content-Disposition", "attachment; filename=test.png")
	//  文件名编码为 二进制
	c.Header("Content-Transfer-Encoding", "binary")
}

func main() {

	router := gin.Default()

	router.GET("/fileDownload", _fileDownload)

	router.Run(":50825")

}

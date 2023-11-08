package main

import (
	logUtil "gin/06-logrus/10"
	"gin/07-integrationLogrus/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	logUtil.InitLogrus("./logFile", "Integration-Logrus")
	router := gin.New()
	router.Use(middleware.MyLogrus())
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"mgs": "你好"})
	})

	router.Run(":80")

}

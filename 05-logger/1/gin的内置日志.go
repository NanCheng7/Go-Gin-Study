package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type LogColor struct {
	Info  string
	Time  string
	Code  string
	Path  string
	Reset string
}

var logColor LogColor = LogColor{
	"\033[1;;32m",
	"\033[1;;36m",
	"\033[1;;35m",
	"\033[1;;31m",
	"\033[0m",
}

func _printfFormat(params gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"[ %sNanCheng - INFO%s ] Time => %s%s%s %s|%d|%s Method ==> Path: %s%-5s%s ==> %s%s%s",
		logColor.Info, logColor.Reset,
		logColor.Time, params.TimeStamp.Format("2006-01-02 15:04:05"), logColor.Reset,
		logColor.Code, params.StatusCode, logColor.Reset,
		params.MethodColor(), params.Method, params.ResetColor(),
		logColor.Path, params.Path, logColor.Reset,
	)
}

func main() {

	//更改Debug的输出格式
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf(
	//		"[ NanCheng - INFO ] %s  %s ===> %s  (%d MiddleWires) \n",
	//		httpMethod,
	//		absolutePath,
	//		handlerName,
	//		nuHandlers,
	//	)
	//}
	//
	//存储到文件
	//file, err := os.Create("gin.log")
	//if err != nil {
	//	defer recover()
	//}
	//gin.DefaultWriter = file
	//gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	gin.SetMode(gin.ReleaseMode)

	//router := gin.Default()
	router := gin.New()

	//router.Use(gin.LoggerWithFormatter(_printfFormat))
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: _printfFormat, Output: os.Stdout}))

	router.GET("/index", func(context *gin.Context) {

	})

	router.Run(":80")
}

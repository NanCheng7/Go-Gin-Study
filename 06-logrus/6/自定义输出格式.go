package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

const (
	Black int = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	Gray
)

type MyFormat struct {
	Prefix string
}

func (mf *MyFormat) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的日志等级换色
	var color int

	switch entry.Level {
	case logrus.ErrorLevel:
		color = Red
	case logrus.WarnLevel:
		color = Yellow
	case logrus.InfoLevel:
		color = Cyan
	default:
		color = 9
	}

	//设置buffer 缓冲区
	var buffer *bytes.Buffer

	if entry.Buffer == nil {
		buffer = &bytes.Buffer{}
	} else {
		buffer = entry.Buffer
	}

	//时间格式化
	formatTime := entry.Time.Format("2006-01-02 15:04:05")
	//具体的函数名
	funcName := entry.Caller.Func.Name()
	//文件的行号
	fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	// 设置格式
	fmt.Fprintf(
		buffer,
		"[ NanCheng <==> %s ]  \u001B[3%dm[ %s ]\u001B[0m  [ time  => %s ]"+
			"\n=>[ caller position  => %s ]  [ caller function => %s ]"+
			"\n=>[ message => %s ] \n",
		mf.Prefix, color, entry.Level, formatTime, fileVal, funcName, entry.Message)
	return buffer.Bytes(), nil
}

func main() {
	logrus.SetFormatter(&MyFormat{"GIN"})
	logrus.SetReportCaller(true) //开启返回函数名和行号
	//路径
	path := "./logFile/DiyInfo.log"
	// 保存文件
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	logrus.Info("NanCheng")
	logrus.Error("NanCheng")
	logrus.Warn("NanCheng")
	logrus.Debug("NanCheng")
}

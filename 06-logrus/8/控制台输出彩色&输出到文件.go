package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
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

//  定义一个日志记录文件的位置
const (
	ErrorLogPath   = "./logFile/error.log"
	DefaultLogPath = "./logFile/logs.log"
)

type MyFormat struct {
	Prefix string
}

// Format 自定义输入输出格式
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

type MyHook struct {
	Formatter logrus.Formatter
}

// Levels 设置哪些日志才会生效
func (mh *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire 设置 Field字段
func (mh *MyHook) Fire(e *logrus.Entry) error {
	//根据hook的Formatter 获取日志信息
	Message, _ := mh.Formatter.Format(e)
	//对错误日志进行分流
	if e.Level == logrus.ErrorLevel {
		//获取Error日志文件的输入流
		errorFile, _ := os.OpenFile(ErrorLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		defer errorFile.Close()
		errorFile.Write(Message)
	} else {
		//获取其他日志文件的输入流
		otherFile, _ := os.OpenFile(DefaultLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		defer otherFile.Close()
		otherFile.Write(Message)
	}
	return nil
}

func main() {
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	logger.SetFormatter(&MyFormat{"GIN"})

	logger.SetReportCaller(true) //开启返回函数名和行号

	hook := &MyHook{
		Formatter: &logrus.TextFormatter{
			DisableColors: true, TimestampFormat: "2006-01-02 15:04:05",
		}}
	logger.AddHook(hook)
	logger.Error("error")
	logger.Info("info")
	logger.Warn("warn")
	logger.Debug("debug")

}

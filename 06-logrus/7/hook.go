package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

type MyHook struct {
}

// Levels 设置哪些日志才会生效
func (mh *MyHook) Levels() []logrus.Level {
	//设置只对error生效
	return []logrus.Level{logrus.ErrorLevel}
}

// Fire 设置 Field字段
func (mg *MyHook) Fire(e *logrus.Entry) error {
	//e.Data["NanCheng"] = "Qianjue"
	file, _ := os.OpenFile("./logFile/error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	allMessage, _ := e.String()
	file.Write([]byte(allMessage))
	return nil
}

func main() {
	logrus.AddHook(&MyHook{})
	logrus.Warnln("你好")
	logrus.Errorf("你好")
}

package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {

	file, _ := os.OpenFile("./logFile/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	logrus.SetFormatter(
		&logrus.TextFormatter{
			ForceColors: true,
		})
	logrus.Warn("警告")
	logrus.Error("出错了")
	logrus.Info("信息")
	logrus.Debug("debug")

}

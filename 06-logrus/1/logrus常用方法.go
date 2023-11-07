package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Error("出错了")
	logrus.Warning("警告")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("debug")
	logrus.Println("print")

	fmt.Println(logrus.GetLevel())
}

package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.WithField("app", "study").
		WithField("NanCheng", "User")
	log.Error("你好")
	log2 := log.WithFields(
		logrus.Fields{
			"app":      "study",
			"NanCheng": "User",
		})
	log2.Info("ba")
}

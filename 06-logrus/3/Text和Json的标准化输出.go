package main

import "github.com/sirupsen/logrus"

func main() {

	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(
		&logrus.TextFormatter{
			ForceColors: true,
		})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Error("南城")
	logrus.Warn("南城")
	logrus.Info("南城")
	logrus.Debug("南城")

}

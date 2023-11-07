package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// LogFormatter 自定以的logFormatter
type LogFormatter struct {
}

// Format 自定义的格式方法
func (lf *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	currentTime := time.Now().Format("2006-01-02")
	//唤起的文件位置和行号
	var fileVal string
	var lineVal int
	if entry.Caller != nil {
		fileVal = filepath.Base(entry.Caller.File)
		lineVal = entry.Caller.Line
	}
	msg := fmt.Sprintf("[ %s ] | time => %s | [position =>  %s : %d] \n=>>Msg: %s \n", strings.ToUpper(entry.Level.String()), currentTime, fileVal, lineVal, entry.Message)
	return []byte(msg), nil
}

type FileLevelHook struct {
	Files map[string]*os.File
}

func (hook *FileLevelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *FileLevelHook) Fire(entry *logrus.Entry) error {
	myFormat := LogFormatter{}
	msg, _ := myFormat.Format(entry)
	// 全局日志记录
	hook.Files["all"].Write(msg)
	// 分离日志
	switch entry.Level {
	case logrus.ErrorLevel:
		hook.Files["error"].Write(msg)
	case logrus.InfoLevel:
		hook.Files["info"].Write(msg)
	case logrus.WarnLevel:
		hook.Files["warn"].Write(msg)
	}
	return nil
}

// InitLogrus 初始化Loggrus
func InitLogrus(FilePath, AppName string) {
	// 拼接文件路径的前缀
	prefix := FilePath + "/" + AppName
	//初始化hook
	hook := FileLevelHook{
		Files: make(map[string]*os.File),
	}
	//创建文件夹 和 hook 对应的文件流
	os.MkdirAll(prefix, os.ModePerm)
	hook.Files["error"], _ = os.OpenFile(prefix+"/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	hook.Files["info"], _ = os.OpenFile(prefix+"/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	hook.Files["warn"], _ = os.OpenFile(prefix+"/warn.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	hook.Files["all"], _ = os.OpenFile(prefix+"/all.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// logrus的初始化设置
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.AddHook(&hook)
}
func main() {
	InitLogrus("./logFile", "NanCheng")
	logrus.Info("NanCheng--info")
	logrus.Infof("NanCheng--infof")
	logrus.Error("NanCheng--error")
	logrus.Errorf("NanCheng--errorf")
	logrus.Warn("NanCheng--warn")
	logrus.Warning("NanCheng--warning")
	logrus.Warnln("NanCheng--warnln")
	logrus.Debug("NanCheng--debug")
	logrus.Debugf("NanCheng--debugf")
}

package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type MyHook struct {
	File     *os.File
	LogPath  string
	FileDate string
	AppName  string
	LF       LogFormatter_Hook
}

func (mh *MyHook) Levels() []log.Level {
	return log.AllLevels
}

// LogFormatter_Hook 自定以的logFormatter
type LogFormatter_Hook struct {
}

// Format 自定义的格式方法
func (lf *LogFormatter_Hook) Format(entry *log.Entry) ([]byte, error) {
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

func (mh *MyHook) Fire(e *log.Entry) error {
	if mh == nil {
		return errors.New("logFileWriter is nil")
	}
	if mh.File == nil {
		return errors.New("file not opened")
	}
	//根据自定义的Format来获取信息
	msg, _ := mh.LF.Format(e)
	currentTime := time.Now().Format("2006-01-02")
	//判断日期是否相等
	if mh.FileDate != currentTime {
		mh.File.Close()
		err := os.MkdirAll(fmt.Sprintf("%s/%s", mh.LogPath, currentTime), os.ModePerm)
		if err != nil {
			return err
		}
		fileName := fmt.Sprintf("%s/%s/%s.log", mh.LogPath, currentTime, mh.AppName)
		mh.File, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
	}
	_, err := mh.File.Write(msg)
	return err
}
func InitLogrus(logPath, appName string) {
	fileDate := time.Now().Format("2006-01-02")
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		log.Error(err)
		return
	}
	filename := fmt.Sprintf("%s/%s/%s.log", logPath, fileDate, appName)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Error(err)
		return
	}
	myFormat := LogFormatter_Hook{}
	myHook := MyHook{
		file,
		logPath,
		fileDate,
		appName,
		myFormat,
	}
	log.AddHook(&myHook)
	log.SetFormatter(&myFormat)
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
}
func main() {
	InitLogrus("./logFile", "NanCheng-HOOK")
	for {
		log.Error("this is a error log" + time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(20 * time.Second)
	}
}

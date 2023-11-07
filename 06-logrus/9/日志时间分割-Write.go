package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// LogFormatter 自定以的logFormatter
type LogFormatter struct {
}

// Format 自定义的格式方法
func (lf *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
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

type LogFileWriter struct {
	File     *os.File
	LogPath  string
	FileDate string
	AppName  string
}

// Write 写入信息到文件方法
func (lfw *LogFileWriter) Write(data []byte) (n int, err error) {
	if lfw == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if lfw.File == nil {
		return 0, errors.New("file not opened")
	}

	//判断是否需要切换日期
	fileDate := time.Now().Format("2006-01-02")
	if lfw.FileDate != fileDate {
		lfw.File.Close()
		err = os.MkdirAll(fmt.Sprintf("%s/%s", lfw.LogPath, fileDate), os.ModePerm)
		if err != nil {
			return 0, err
		}
		filename := fmt.Sprintf("%s/%s/%s-%s.log", lfw.LogPath, fileDate, lfw.AppName, fileDate)

		lfw.File, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
		if err != nil {
			return 0, err
		}
	}
	n, e := lfw.File.Write(data)
	return n, e
}

func InitLog(logPath string, appName string) {
	fileDate := time.Now().Format("2006-01-02")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		log.Error(err)
		return
	}

	filename := fmt.Sprintf("%s/%s/%s-%s.log", logPath, fileDate, appName, fileDate)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Error(err)
		return
	}

	fileWriter := LogFileWriter{file, logPath, fileDate, appName}

	log.SetOutput(os.Stdout)
	writers := []io.Writer{
		&fileWriter,
		os.Stdout}
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		log.SetOutput(fileAndStdoutWriter)
	} else {
		log.Info("failed to log to file.")
	}
	log.SetReportCaller(true)
	log.SetFormatter(&LogFormatter{})
}
func main() {
	InitLog("./logFile", "NanCheng")
	log.Info("NanCheng1")
	log.Error("NanCheng2")
	log.Warn("NanCheng3")
}

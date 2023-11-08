package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// 每一次响应携带的参数
type LogParams struct {
	Request *http.Request

	// StatusCode is HTTP response code.
	StatusCode int
	// Latency is how much time the server cost to process a certain request.
	TakeTime time.Duration
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string
	// isTerm shows whether gin's output descriptor refers to a terminal.
	isTerm bool
	// BodySize is the size of the Response Body
	BodySize int
	// Keys are the keys set on the request's context.
	Keys map[string]any
}

const (
	status200 = 32
	status404 = 33
	status500 = 31

	methodGET = 34
)

// MyLogrus 中间件
func MyLogrus() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("this is myLogrus MiddleWare")
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		startTime := time.Now()
		logParams := LogParams{
			Request: c.Request,
			Keys:    c.Keys,
			isTerm:  true,
		}
		c.Next()
		//时间差
		logParams.TakeTime = time.Now().Sub(startTime)

		logParams.ClientIP = c.ClientIP()
		logParams.Method = c.Request.Method
		logParams.StatusCode = c.Writer.Status()
		logParams.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		logParams.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		var statusColor string
		switch logParams.StatusCode {
		case 200:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", status200, logParams.StatusCode)
		case 404:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", status404, logParams.StatusCode)
		}
		var methodColor string
		switch logParams.Method {
		case "GET":
			methodColor = fmt.Sprintf("\033[%dm %s \033[0m", methodGET, logParams.Method)

		}

		logrus.Infof("[GIN] |%s| %vms | %s | %s | %s |  Key => %s ",
			statusColor, logParams.TakeTime.Milliseconds(), logParams.ClientIP, methodColor, path, logParams.Keys,
		)

	}
}

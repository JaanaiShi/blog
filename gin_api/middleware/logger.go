package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {

	// 将日志输入到文件中
	filePath := "log/log"
	linkName := "latest_log.log"   // 建立软链接
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}

	logger := logrus.New()
	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		retalog.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})


	logger.AddHook(Hook)



	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()   // 中间件运作方式：洋葱模型
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()) / 1000000.0)))
		fmt.Println(spendTime)
		hostName, err  := os.Hostname()
		if err != nil {
			hostName = "unknow"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()     // 查看客户端的IP
		userAgent := c.Request.UserAgent()  // 客户端信息
		dataSize := c.Writer.Size()   // 长度
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method     // 请求方法
		path := c.Request.RequestURI     // 路径

		entry := logger.WithFields(logrus.Fields{
			"StartTime": startTime,
			"HostName": hostName,
			"status": statusCode,
			"SpendTime": spendTime,
			"IP": clientIp,
			"Method": method,
			"Path": path,
			"dataSize":dataSize,
			"Agent":userAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		}else if statusCode >= 400 {
			entry.Warn()
		}else {
			entry.Info()
		}
	}
}

package base

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

func init() {
	// 定义日志格式
	//format := &log.TextFormatter{}
	format := &prefixed.TextFormatter{}
	format.ForceFormatting = true
	log.SetFormatter(format)

	// 设置日志级别，默认是 info
	environment := os.Getenv("environment")
	if "production" != environment {
		log.SetLevel(log.DebugLevel)
	}

	// 开启高亮显示
	format.DisableColors = false
	format.ForceColors = true

	// 开启时间戳
	format.FullTimestamp = true
	// https://www.zhihu.com/question/366830553
	// 000000 表示精确到毫秒
	format.TimestampFormat = "2006-01-02 15:04:05.000000"
	//logFile := kvs.GetCurrentFilePath("access_log.%Y%m%d",1)
	//out, _ := rotatelogs.New(logFile)
	//log.SetOutput(out)
}

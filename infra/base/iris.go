package base

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	irisRecover "github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
	"resk/infra"
	"time"
)

var irisApplication *iris.Application

func Iris() *iris.Application {
	return irisApplication
}

type IrisStarter struct {
	infra.BaseStarter
}

func (s *IrisStarter) Init(ctx infra.StarterContext) {
	// 创建 irisApplication 实例
	irisApplication = initIrisApplication()
	// 日志统一，使用 logrus
	l := irisApplication.Logger()
	l.Install(logrus.StandardLogger())
}

func (s *IrisStarter) Start(ctx infra.StarterContext) {
	// 把路由信息打印出来
	for _, v := range irisApplication.GetRoutes() {
		logrus.Info(v.Trace())
	}
	// 启动iris框架
	port := ctx.Props().GetDefault("app.service.port", "8888")
	irisApplication.Run(iris.Addr(":" + port))

}

func (s *IrisStarter) StartBlocking() bool {
	return true
}

func initIrisApplication() *iris.Application {
	app := iris.New()

	// 中间件
	app.Use(irisRecover.New())

	cfg := logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		// 日志格式定义
		LogFunc: func(endTime time.Time,
			latency time.Duration,
			status, ip, method, path string,
			message interface{},
			headerMessage interface{},
		) {
			app.Logger().Infof("| %s | %s | %s | %s | %s | %s | %s | %s |",
				endTime.Format("2006-01-02 15:04.05.000000"),
				latency.String(),
				status,
				ip,
				method,
				path,
				message,
				headerMessage,
			)
		},
	}

	// 中间件
	app.Use(logger.New(cfg))
	return app
}

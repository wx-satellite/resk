package infra

import "github.com/tietang/props/kvs"

type BootApplication struct {
	conf kvs.ConfigSource
	starterContext StarterContext
}


func New(conf kvs.ConfigSource) *BootApplication {
	app := &BootApplication{
		conf:conf,
		starterContext:StarterContext{},
	}
	app.starterContext[PropsKey] = conf
	return app
}


func (b *BootApplication) Start() {
	// 初始化starter
	b.init()
	// 安装starter
	b.setUp()
	// 启动starter
	b.start()
}

func (b *BootApplication)  init() {
	for _, v := range StarterRegister.AllStarters() {
		v.Init(b.starterContext)
	}
}


func (b *BootApplication) setUp() {
	for _, v := range StarterRegister.AllStarters() {
		v.Setup(b.starterContext)
	}
}


func (b *BootApplication) start() {
	for k, v := range StarterRegister.AllStarters() {
		// 如果启动会阻塞并且当前不是最后一个则使用 goroutine 异步启动
		if v.StartBlocking() && k + 1 < len(StarterRegister.AllStarters()) {
			go v.Start(b.starterContext)
		}else {
			v.Start(b.starterContext)
		}
	}
}
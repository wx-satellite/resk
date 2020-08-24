package testx

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"resk/infra"
	"resk/infra/base"
)

// 测试初始化
func init() {

	configPath := kvs.GetCurrentFilePath("../brun/config.ini", 1)

	conf := ini.NewIniFileConfigSource(configPath)

	app := infra.New(conf)

	infra.Register(&base.PropsStarter{})

	//infra.Register(&base.DbxDatabaseStarter{})

	infra.Register(&base.ValidatorStarter{})

	app.Start()
}

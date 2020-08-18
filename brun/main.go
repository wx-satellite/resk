package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "resk"
	"resk/infra"
)

func main() {
	configPath := kvs.GetCurrentFilePath("config.ini",1)

	conf := ini.NewIniFileConfigSource(configPath)

	app := infra.New(conf)

	app.Start()
}

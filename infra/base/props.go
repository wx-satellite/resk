package base

import (
	"github.com/tietang/props/kvs"
	"resk/infra"
)

var props kvs.ConfigSource

type PropsStarter struct {
	infra.BaseStarter
}

func Props() kvs.ConfigSource {
	return props
}


func (s *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
}







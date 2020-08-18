package resk

import (
	"resk/infra"
	"resk/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
}

package msa

import (
	"github.com/karincake/apem/appa"
)

type MsConf struct {
	Dsn  string
	Host string
	Port int
}

type MsItf interface {
	Init(*MsConf, *appa.AppConf)
}

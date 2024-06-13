package msa

import (
	"github.com/karincake/apem/appa"
)

type MsCfg struct {
	Dsn  string
	Host string
	Port int
}

type MsItf interface {
	Init(*MsCfg, *appa.AppCfg)
}

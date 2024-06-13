package dba

import (
	"github.com/karincake/apem/appa"
)

type DbCfg struct {
	Dsn          string
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxIdleTime  string `yaml:"maxIdleTime"`
	Dialect      string `yaml:"dialect"`
}

type DbItf interface {
	Init(*DbCfg, *appa.AppCfg)
}

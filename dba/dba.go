package dba

import (
	"github.com/karincake/apem/appa"
)

type DbCfg struct {
	Name         string `yaml:"name"`
	Dsn          string `yaml:"dsn"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxIdleTime  string `yaml:"maxIdleTime"`
	Dialect      string `yaml:"dialect"`
}

type MultiDbCfg struct {
	Dbs []DbCfg `yaml:"dbs"`
}

type DbItf interface {
	Init(*DbCfg, *appa.AppCfg)
	InitMulti(*MultiDbCfg, *appa.AppCfg)
}

package langa

import (
	"github.com/karincake/apem/appa"
)

type LangConf struct {
	Active   string
	Path     string
	FileName string `yaml:"fileName"`
}

type LangItf interface {
	Init(*LangConf, *appa.AppConf)
}

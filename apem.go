package apem

import (
	"log"
	"net/http"

	"github.com/karincake/apem/appa"
	"github.com/karincake/apem/dba"
	"github.com/karincake/apem/httpa"
	"github.com/karincake/apem/langa"
	"github.com/karincake/apem/loggera"
	"github.com/karincake/apem/msa"

	hs "github.com/karincake/apem/http-std"
)

var App *apemConf
var CfgFile string

// init
func init() {
	App = &apemConf{
		Conf:       &appa.AppConf{},
		LoggerConf: &loggera.LoggerConf{},
		LangConf:   &langa.LangConf{},
		DbConf:     &dba.DbConf{},
		MsConf:     &msa.MsConf{},
		HttpConf:   &httpa.HttpConf{},
	}
	App.initConfig()
}

// app start the App
func Run(h http.Handler, m ...any) {
	readines := 0
	loggerIdx := -1

	for i := range m {
		if myModule, ok := m[i].(loggera.LoggerItf); ok {
			readines++
			loggerIdx = i
			myModule.Init(App.LoggerConf, App.Conf)
		} else if myModule, ok := m[i].(langa.LangItf); ok {
			myModule.Init(App.LangConf, App.Conf)
		} else if myModule, ok := m[i].(dba.DbItf); ok {
			myModule.Init(App.DbConf, App.Conf)
		} else if myModule, ok := m[i].(msa.MsItf); ok {
			myModule.Init(App.MsConf, App.Conf)
		}
	}

	if readines < 1 {
		log.Fatal("App.Run doesn't supply enough options. Please make sure implementation of `loggera` is supplied.")
	}

	App.initExtCall()

	hs.O.Init(App.HttpConf, &h, App.Conf, m[loggerIdx].(loggera.LoggerItf))
}

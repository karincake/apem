package apem

import (
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/karincake/apem/appa"
	"github.com/karincake/apem/dba"
	"github.com/karincake/apem/httpa"
	"github.com/karincake/apem/loggera"
	"github.com/karincake/apem/msa"
	"gopkg.in/yaml.v3"

	hs "github.com/karincake/apem/http-std"
)

type extCall func()

var extCalls []extCall

var App *apemCfg
var CfgFile string

// init
func init() {
	App = &apemCfg{
		AppCfg: &appa.AppCfg{
			CodeName: "apem",
			FullName: "Apem Instance",
			Version:  "0.0.1",
			Env:      "development",
		},
		LoggerCfg: &loggera.LoggerCfg{},
		DbCfg:     &dba.DbCfg{},
		MsCfg:     &msa.MsCfg{},
		HttpCfg:   &httpa.HttpCfg{},
	}
	App.initCfg()
}

// app start the App
func Run(h http.Handler, m ...any) {
	readines := 0
	loggerIdx := -1

	initExtCall()

	for i := range m {
		if myModule, ok := m[i].(loggera.LoggerItf); ok {
			readines++
			loggerIdx = i
			myModule.Init(App.LoggerCfg, App.AppCfg)
		} else if myModule, ok := m[i].(dba.DbItf); ok {
			myModule.Init(App.DbCfg, App.AppCfg)
		} else if myModule, ok := m[i].(msa.MsItf); ok {
			myModule.Init(App.MsCfg, App.AppCfg)
		}
	}

	if readines < 1 {
		log.Fatal("Please make sure App.Run supplied by mandatory adapters implementation. Missing: `loggera`.")
	}

	hs.O.Init(App.HttpCfg, &h, App.AppCfg, m[loggerIdx].(loggera.LoggerItf))
}

func ParseCfg(cfg any) {
	yamlFile, err := os.ReadFile(CfgFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	oriVal := reflect.ValueOf(cfg).Interface()
	err = yaml.Unmarshal(yamlFile, oriVal)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func RegisterExtCall(funcToCall extCall) {
	extCalls = append(extCalls, funcToCall)
}

func initExtCall() {
	log.Print("Executing extra calls")
	for _, init := range extCalls {
		init()
	}
}

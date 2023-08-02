package apem

import (
	"net/http"

	db "github.com/karincake/apem/databasegorm"
	lg "github.com/karincake/apem/lang"
	lz "github.com/karincake/apem/loggerzap"
)

type app struct {
	CodeName   string
	FullName   string
	Env        string
	Version    string
	LoggerConf *lz.LoggerConf
	LangConf   *lg.LangConf
	DbConf     *db.DbConf
	HttpConf   *httpConf
}

// export package vars
var Apem *app

// init
func init() {
	Apem = &app{
		LoggerConf: &lz.LoggerConf{},
		LangConf:   &lg.LangConf{},
		DbConf:     &db.DbConf{},
	}
	Apem.initConfig()
}

// app starter
func Run(appCodeName string, routerIn http.Handler) {
	// basic instance completion
	Apem.CodeName = appCodeName
	// fmt.Println(Apem)

	// Call manually to make it goes according to the desired flow
	lz.Init(*Apem.LoggerConf)
	lg.Init(*Apem.LangConf)
	Apem.initExtCall()
	Apem.initHttp(routerIn)
}

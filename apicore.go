package apem

import (
	"net/http"

	db "github.com/karincake/apem/databasegorm"
	lg "github.com/karincake/apem/lang"
	lz "github.com/karincake/apem/loggerzap"
	mr "github.com/karincake/apem/memstorageredis"
)

type app struct {
	CodeName        string
	FullName        string
	Env             string
	Version         string
	LoggerConf      *lz.LoggerConf
	LangConf        *lg.LangConf
	DbConf          *db.DbConf
	MsConf          *mr.MsConf
	HttpConf        *httpConf
	RateLimiterConf *rateLimiterConf
}

// export package vars
var Apem *app

// init
func init() {
	Apem = &app{
		LoggerConf: &lz.LoggerConf{},
		LangConf:   &lg.LangConf{},
		DbConf:     &db.DbConf{},
		MsConf:     &mr.MsConf{},
	}
	Apem.initConfig()
}

// app starter
func Run(appCodeName string, routerIn http.Handler) {
	// basic instance completion
	Apem.CodeName = appCodeName

	// Call manually to make it goes according to the desired flow
	lz.Init(*Apem.LoggerConf)
	lg.Init(*Apem.LangConf)
	db.Init(*Apem.DbConf)
	mr.Init(*Apem.MsConf)
	Apem.initExtCall()
	Apem.initHttp(routerIn)
}

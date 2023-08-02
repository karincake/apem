package apem

import (
	"net/http"
)

type app struct {
	CodeName string
	FullName string
	Env      string
	Version  string
	HttpConf *httpConf
}

// export package vars
var Apem *app

// init
func init() {
	Apem = &app{}
	Apem.initConfig()
}

// app starter
func Run(appCodeName string, routerIn http.Handler) {
	// basic instance completion
	Apem.CodeName = appCodeName
	// fmt.Println(Apem)

	// Call manually to make it goes according to the desired flow
	Apem.initExtCall()
	Apem.initHttp(routerIn)
}

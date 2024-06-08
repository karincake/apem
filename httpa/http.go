package httpa

import (
	"net/http"

	"github.com/karincake/apem/appa"
)

type HttpConf struct {
	Host string
	Port int
}

type RateLimiterConf struct {
	Enabled bool
	Rps     float64
	Burst   int
}

type HttpItf interface {
	Init(*HttpConf, *http.Handler, *appa.AppConf)
}

package httpa

import (
	"net/http"

	"github.com/karincake/apem/appa"
)

type HttpCfg struct {
	Host string
	Port int
}

type RateLimiterCfg struct {
	Enabled bool
	Rps     float64
	Burst   int
}

type HttpItf interface {
	Init(*HttpCfg, *http.Handler, *appa.AppCfg)
}

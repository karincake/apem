package httpstdloggermw

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"time"

	l "github.com/karincake/apem/loggera"
)

var Logger l.LoggerItf

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (obj *wrappedWriter) WriteHeader(statusCode int) {
	obj.statusCode = statusCode
	obj.ResponseWriter.WriteHeader(statusCode)
}

func (obj *wrappedWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hj, ok := obj.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("the hijacker interface is not supported")
	}

	return hj.Hijack()
}

func (obj *wrappedWriter) Flush() {
	if fl, ok := obj.ResponseWriter.(http.Flusher); ok {
		if obj.statusCode == 0 {
			obj.WriteHeader(http.StatusOK)
		}

		fl.Flush()
	}
}

func SetLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		Logger.Info().
			String("scope", "request").
			Int("status", wrapped.statusCode).
			String("method", r.Method).
			String("path", r.URL.Path).
			String("query", r.URL.RawQuery).
			String("duration", time.Since(start).String()).Send()
	})
}

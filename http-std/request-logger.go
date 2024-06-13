package httpstd

import (
	"net/http"
	"time"

	l "github.com/karincake/apem/loggera"
)

var logger l.LoggerItf

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (obj *wrappedWriter) WriteHeader(statusCode int) {
	obj.statusCode = statusCode
	obj.ResponseWriter.WriteHeader(statusCode)
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		logger.Info().
			String("scope", "request").
			Int("status", wrapped.statusCode).
			String("method", r.Method).
			String("path", r.URL.Path).
			String("duration", time.Since(start).String()).Send()
	})
}

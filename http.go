package apem

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"

	lz "github.com/karincake/apem/loggerzap"
)

type httpConf struct {
	Host     string
	Port     int
	HttpConf *httpConf
}

var wg sync.WaitGroup

func (a *app) initHttp(handler http.Handler) {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", a.HttpConf.Host, a.HttpConf.Port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		s := <-sig
		lz.I.Info("process", zap.String("type", "signal"), zap.String("source", "signal"), zap.String("act", s.String()), zap.String("status", "delegated"))
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}
		lz.I.Info("process", zap.String("type", "background"), zap.String("source", "httprouter"), zap.String("act", "closing"), zap.String("addr", srv.Addr), zap.String("status", "done"))

		wg.Wait()
		shutdownError <- nil
	}()

	lz.I.Info("process", zap.String("type", "server"), zap.String("source", "httprouter"), zap.String("act", "serve"), zap.String("addr", srv.Addr), zap.String("status", "running"))
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		lz.I.Fatal(err.Error())
	}

	err = <-shutdownError
	if err != nil {
		lz.I.Fatal(err.Error())
	}

	lz.I.Info("process", zap.String("type", "server"), zap.String("source", "httprouter"), zap.String("act", "shutdown"), zap.String("addr", srv.Addr), zap.String("status", "done"))
}

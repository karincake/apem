package httpstd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/karincake/apem/appa"
	"github.com/karincake/apem/httpa"
	"github.com/karincake/apem/loggera"
)

type httpStd struct{}

var O httpStd = httpStd{}
var wg sync.WaitGroup

func (o *httpStd) Init(c *httpa.HttpConf, h *http.Handler, a *appa.AppConf, l loggera.LoggerItf) {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", c.Host, c.Port),
		Handler:      *h,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Gracefull shutdown
	shutdownError := make(chan error)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		s := <-sig

		l.Info().String("source", "net/http").String("action", "stopage").String("status", "done").String("signal", s.String()).String("host", c.Host).Int("port", c.Port).Send()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		wg.Wait()
		shutdownError <- nil
	}()

	l.Info().String("source", "net/http").String("action", "instantiation").String("status", "listening").String("host", c.Host).Int("port", c.Port).Send()
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err.Error())
	}

	err = <-shutdownError
	if err != nil {
		log.Fatal(err.Error())
	}
	l.Info().String("source", "net/http").String("action", "stopage").String("status", "done").String("host", c.Host).Int("port", c.Port).Send()
}

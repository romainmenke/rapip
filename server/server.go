package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/romainmenke/rapip/ratelimiter"
	"github.com/romainmenke/rapip/router"
)

type Config struct {
	Port string
}

func Run(config Config) {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(
		signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	// handler
	handler := http.Handler(router.New())

	// rate limiter
	ratelimiter := ratelimiter.NewLimiter(1000, 50)
	handler = ratelimiter.Handler(handler)

	server := &http.Server{
		Addr:              ":" + config.Port,
		Handler:           handler,
		IdleTimeout:       20 * time.Second,
		MaxHeaderBytes:    1 << 16,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      120 * time.Second,
	}

	go func() {
		<-signal_chan

		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		defer cancel()

		server.Shutdown(ctx)
	}()

	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		return
	}
	if err != nil {
		panic(err)
	}
}

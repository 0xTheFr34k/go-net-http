package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/0xTheFr34k/go-net-http/middleware"
	"github.com/0xTheFr34k/go-net-http/router"
)

func main() {
	appMux := http.NewServeMux()
	appMux.Handle("/v1/", http.StripPrefix("/v1", router.V1Router()))

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Addr:    ":8000",
		Handler: appMux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err)
			return
		}
	}()

	<-shutdown.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		middleware.Wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-ctx.Done():
	}

	server.Shutdown(context.Background())
}

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	apirest "github.com/alonsojl/matryer"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type config struct {
	Host string
	Port string
}

func run(ctx context.Context, w io.Writer, args []string) error {
	var (
		config = config{
			Host: os.Getenv("Host"),
			Port: os.Getenv("Port"),
		}
		logger = apirest.NewLogger()
	)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	srv := newServer(logger)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}

	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		// make a new context for the Shutdown (thanks Alessandro Rosetti)
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func newServer(logger *logrus.Logger) http.Handler {
	mux := chi.NewMux()
	addRoutes(mux, logger)
	var handler http.Handler = mux
	handler = someMiddleware(handler)
	handler = adminOnly(handler)
	return handler
}

func addRoutes(mux *chi.Mux, logger *logrus.Logger) {
	mux.Get("/healthz", handleSomething(logger))
}

func handleSomething(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Println("example handle")
	}
}

func someMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("ENV") != "development" {
			w.WriteHeader(http.StatusNotFound)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func adminOnly(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAdmin := true
		if !isAdmin {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

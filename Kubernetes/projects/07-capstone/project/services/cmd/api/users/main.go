package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type (
	Server interface {
		Serve() error
		Shutdown(ctx context.Context) error
	}

	server struct {
		ctx        context.Context
		httpServer *http.Server
	}
)

// addRoutes is where the entire API surface is mapped
// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo
func addRoutes(mux *http.ServeMux) {
	mux.Handle("/*", http.NotFoundHandler())
}

// Heartbeat logs incoming requests on global HTTP handler
func heartbeat(next http.Handler, endpoint string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == "GET" || r.Method == "HEAD") && strings.EqualFold(r.URL.Path, endpoint) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("."))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	addRoutes(mux)

	var handler http.Handler = mux
	handler = heartbeat(handler, "/health")

	return handler
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *server) Serve() error {
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func newServer(ctx context.Context, handler http.Handler) Server {
	const (
		maxTimeout   = 120 * time.Second
		readTimeout  = 5 * time.Second
		writeTimeout = 2 * readTimeout
		idleTimeout  = maxTimeout
	)

	httpserver := &http.Server{
		Addr:         net.JoinHostPort("0.0.0.0", strconv.Itoa(8080)),
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	srv := &server{
		ctx:        ctx,
		httpServer: httpserver,
	}

	return srv
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	infoLog := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
	errLog := log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
	router := NewRouter()
	srv := newServer(ctx, router)

	var serveError error

	go func() {
		if err := srv.Serve(); err != nil {
			serveError = err
		}
	}()

	infoLog.Print("http://0.0.0.0:8080/health")

	if serveError != nil {
		return serveError
	}
	// Wait for shutdown signal
	<-ctx.Done()
	infoLog.Print("shutdown signal received")

	// Allow some time for graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		errLog.Print(err, "server shutdown failed")
		return err
	}

	return nil
}

func main() {
	rootCtx := context.Background()

	if err := run(
		rootCtx,
	); err != nil {
		log.Fatalf("error: %+v", err)
	}
}

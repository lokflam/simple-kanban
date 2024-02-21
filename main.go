package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lokflam/simple-kanban/internal/health"
	"github.com/lokflam/simple-kanban/internal/kanban"
)

func main() {
	const Timeout = 20 * time.Second

	// get port from env
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	// configure dependencies
	requestLogger := httplog.NewLogger("request-logger", httplog.Options{
		// JSON:          true,
		TimeFieldName: "time",
	})

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(fmt.Errorf("unable to create connection pool: %w", err))
	}
	defer dbpool.Close()

	// create router
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(requestLogger))
	r.Use(middleware.Timeout(Timeout))
	r.Get("/health", health.Handler())
	r.Mount("/", kanban.Router(dbpool))

	// start server
	server := &http.Server{Addr: ":" + port, Handler: r}
	go func() {
		requestLogger.Logger.Info(fmt.Sprintf("starting server on port %v.", port))
		err := server.ListenAndServe()

		// ignore returned err when server is closed expectedly
		if err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("server failed: %w", err))
		}
	}()

	// catch graceful shutdown signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)

	<-sigCh

	// graceful shutdown server
	gracefulTimeout := Timeout + time.Second
	requestLogger.Logger.Info(fmt.Sprintf("gracefully shutting down server with timeout %v seconds.", gracefulTimeout.Seconds()))
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), gracefulTimeout)
	defer cancelShutdown()

	err = server.Shutdown(shutdownCtx)
	if err != nil {
		panic(fmt.Errorf("failed to graceful shutdown server: %w", err))
	}
	requestLogger.Logger.Info("server closed.")
}

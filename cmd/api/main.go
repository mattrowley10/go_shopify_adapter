package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mattrowley10/the_faywood_adapter/api"
	"github.com/mattrowley10/the_faywood_adapter/internal/config"
)

type App struct {
	logger *slog.Logger
	mux    *http.ServeMux
}

func (a *App) Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	a.logger = logger
	logger.Info("loading config...")

	cfg, err := config.LoadEnv()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	logger.Info("config loaded successfully", "cfg", cfg)

	logger.Info("starting server...")

	server := api.NewServer(logger)

	httpserver := &http.Server{
		Addr:    cfg.Server.ListenAddr,
		Handler: server.Router(),
	}
	errChan := make(chan error, 1)

	go func() {
		logger.Info("listening on address: ", "addr", cfg.Server.ListenAddr)
		errChan <- httpserver.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		logger.Error("error returned on server start", "err", err)
		os.Exit(1)
	case <-stop:
		logger.Info("gracefully shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		httpserver.Shutdown(ctx)
		logger.Info("server shutdown successfully")
	}
}

func main() {
	app := &App{}
	app.Run()
}

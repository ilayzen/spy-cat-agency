package main

import (
	"flag"
	"log"
	"net/http"

	delivery "github.com/ilayzen/spy-cat-agency/internal/delivery/rest"
	"github.com/ilayzen/spy-cat-agency/internal/repository"
	"github.com/ilayzen/spy-cat-agency/internal/service"
	"github.com/ilayzen/spy-cat-agency/pkg/config"
	postgres "github.com/ilayzen/spy-cat-agency/pkg/database"
	"github.com/ilayzen/spy-cat-agency/pkg/logger"
)

const defaultConfigPath = "./cmd/config.yaml"

func main() {
	var configPath = flag.String(
		"config",
		defaultConfigPath,
		"path to config file",
	)

	flag.Parse()

	cfg := Config{}
	err := config.LoadFromFile(*configPath, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	log := logger.NewLogger(cfg.Logging).WithField("app", "main")
	mainLogger := log.WithField("component", "main")

	mainLogger.Infof("config.go: %+v, loaded", cfg)

	db, err := postgres.NewDB(cfg.DB, log)
	if err != nil {
		mainLogger.Fatalf("cannot create new datebase, err: %v", err)
	}

	repos := repository.NewRepository(db)
	svc := service.NewService(log, repos)

	handler := delivery.NewHandler(svc)

	router := delivery.NewRouter(handler)

	srv := &http.Server{
		Addr:         cfg.API.ServerPort,
		WriteTimeout: cfg.API.RequestRWTimeout,
		ReadTimeout:  cfg.API.RequestRWTimeout,
		IdleTimeout:  cfg.API.IdleTimeout,
		Handler:      router,
	}

	// todo: add graceful shutdown
	if err = srv.ListenAndServe(); err != nil {
		mainLogger.Fatalf("cannot start http server, err: %v", err)

	}
}

package main

import (
	"time"

	postgres "github.com/ilayzen/spy-cat-agency/pkg/database"

	"github.com/ilayzen/spy-cat-agency/pkg/logger"
)

type Config struct {
	Logging logger.Config     `yaml:"logging"`
	API     APIConfig         `yaml:"api"`
	DB      postgres.DBConfig `yaml:"db"`
}

type APIConfig struct {
	ServerPort       string        `yaml:"server_port"`
	RequestRWTimeout time.Duration `yaml:"request_rw_timeout"`
	IdleTimeout      time.Duration `yaml:"idle_timeout"`
}
